package manager

import (
	"context"
	"fmt"

	gfnt "github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/kris-nova/logger"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/cfn/outputs"
	"github.com/weaveworks/eksctl/pkg/vpc"
)

// FixClusterCompatibility adds any resources missing in the CloudFormation stack in order to support new features
// like Managed Nodegroups and Fargate
func (c *StackCollection) FixClusterCompatibility(ctx context.Context) error {
	logger.Info("checking cluster stack for missing resources")
	stack, err := c.DescribeClusterStackIfExists(ctx)
	if err != nil {
		return err
	}
	if stack == nil {
		return &StackNotFoundErr{ClusterName: c.spec.Metadata.Name}
	}

	var (
		clusterDefaultSG string
		fargateRole      string
	)

	featureOutputs := map[string]outputs.Collector{
		// available on clusters created after Managed Nodes support was out
		outputs.ClusterDefaultSecurityGroup: func(v string) error {
			clusterDefaultSG = v
			return nil
		},
		// available on 1.14 clusters created after Fargate support was out
		outputs.FargatePodExecutionRoleARN: func(v string) error {
			fargateRole = v
			return nil
		},
	}

	if err := outputs.Collect(*stack, nil, featureOutputs); err != nil {
		return err
	}

	stackSupportsManagedNodes := false
	if clusterDefaultSG != "" {
		stackSupportsManagedNodes, err = c.hasManagedToUnmanagedSG(ctx)
		if err != nil {
			return err
		}
	}

	managedNodeUpdateRequired := !stackSupportsManagedNodes && len(c.spec.ManagedNodeGroups) > 0

	stackSupportsFargate := fargateRole != ""
	fargateUpdateRequired := !stackSupportsFargate && len(c.spec.FargateProfiles) > 0

	if !managedNodeUpdateRequired && !fargateUpdateRequired {
		logger.Info("cluster stack has all required resources")
		return nil
	}

	if managedNodeUpdateRequired {
		logger.Info("cluster stack is missing resources for Managed Nodegroups")
	}
	if fargateUpdateRequired {
		logger.Info("cluster stack is missing resources for Fargate")
	}

	logger.Info("adding missing resources to cluster stack")
	_, err = c.AppendNewClusterStackResource(ctx, false, false)
	return err
}

func (c *StackCollection) hasManagedToUnmanagedSG(ctx context.Context) (bool, error) {
	stackTemplate, err := c.GetStackTemplate(ctx, c.MakeClusterStackName())
	if err != nil {
		return false, fmt.Errorf("error getting cluster stack template: %w", err)
	}
	stackResources := gjson.Get(stackTemplate, resourcesRootPath)
	return builder.HasManagedNodesSG(&stackResources), nil
}

// EnsureMapPublicIPOnLaunchEnabled sets this subnet property to true when it is not set or is set to false
func (c *StackCollection) EnsureMapPublicIPOnLaunchEnabled(ctx context.Context) error {
	// First, make sure we enable the options in EC2. This is to make sure the settings are applied even
	// if the stacks in Cloudformation have the setting enabled (since a stack update would produce "nothing to change"
	// and therefore the setting would not be updated)
	publicIDs := c.spec.VPC.Subnets.Public.WithIDs()
	logger.Debug("enabling attribute MapPublicIpOnLaunch via EC2 on subnets %q", publicIDs)
	err := vpc.EnsureMapPublicIPOnLaunchEnabled(ctx, c.ec2API, publicIDs)
	if err != nil {
		return err
	}

	// Get stack template
	stackName := c.MakeClusterStackName()
	currentTemplate, err := c.GetStackTemplate(ctx, stackName)
	if err != nil {
		return fmt.Errorf("unable to retrieve cluster stack %q: %w", stackName, err)
	}

	// Find subnets in stack
	outputTemplate := gjson.Get(currentTemplate, outputsRootPath)
	publicSubnetsNames, err := getPublicSubnetResourceNames(outputTemplate.Raw)
	if err != nil {
		// Subnets do not appear in the stack because the VPC was imported
		logger.Debug(err.Error())
		return nil
	}

	// Modify the subnets' properties in the stack
	logger.Debug("ensuring subnets have MapPublicIpOnLaunch enabled")
	for _, subnet := range publicSubnetsNames {
		path := subnetResourcePath(subnet)

		currentValue := gjson.Get(currentTemplate, path)
		if !currentValue.Exists() || !currentValue.Bool() {
			currentTemplate, err = sjson.Set(currentTemplate, path, gfnt.True())
			if err != nil {
				return fmt.Errorf("unable to set MapPublicIpOnLaunch property on subnet %q: %w", path, err)
			}
		}
	}
	description := fmt.Sprintf("update public subnets %q with property MapPublicIpOnLaunch enabled", publicSubnetsNames)
	if err := c.UpdateStack(ctx, UpdateStackOptions{
		StackName:     stackName,
		ChangeSetName: c.MakeChangeSetName("update-subnets"),
		Description:   description,
		TemplateData:  TemplateBody(currentTemplate),
		Wait:          true,
	}); err != nil {
		return fmt.Errorf("unable to update subnets: %w", err)
	}
	return nil
}

func subnetResourcePath(subnetName string) string {
	return fmt.Sprintf("Resources.%s.Properties.MapPublicIpOnLaunch", subnetName)
}

// getPublicSubnetResourceNames returns the stack resource names for the public subnets, gotten from the stack
// output "SubnetsPublic"
func getPublicSubnetResourceNames(outputsTemplate string) ([]string, error) {
	publicSubnets := gjson.Get(outputsTemplate, "SubnetsPublic.Value.Fn::Join.1.#.Ref")
	if !publicSubnets.Exists() || len(publicSubnets.Array()) == 0 {
		subnetsJSON := gjson.Get(outputsTemplate, "SubnetsPublic.Value")
		return nil, fmt.Errorf("resource name for public subnets not found. Found %q", subnetsJSON.Raw)
	}
	subnetStackNames := make([]string, 0)
	for _, subnet := range publicSubnets.Array() {
		subnetStackNames = append(subnetStackNames, subnet.String())
	}
	return subnetStackNames, nil
}
