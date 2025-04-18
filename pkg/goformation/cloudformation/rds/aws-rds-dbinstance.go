package rds

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/cloudformation"
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// DBInstance AWS CloudFormation Resource (AWS::RDS::DBInstance)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html
type DBInstance struct {

	// AllocatedStorage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-allocatedstorage
	AllocatedStorage *types.Value `json:"AllocatedStorage,omitempty"`

	// AllowMajorVersionUpgrade AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-allowmajorversionupgrade
	AllowMajorVersionUpgrade *types.Value `json:"AllowMajorVersionUpgrade,omitempty"`

	// ApplyImmediately AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-applyimmediately
	ApplyImmediately *types.Value `json:"ApplyImmediately,omitempty"`

	// AssociatedRoles AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-associatedroles
	AssociatedRoles []DBInstance_DBInstanceRole `json:"AssociatedRoles,omitempty"`

	// AutoMinorVersionUpgrade AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-autominorversionupgrade
	AutoMinorVersionUpgrade *types.Value `json:"AutoMinorVersionUpgrade,omitempty"`

	// AutomaticBackupReplicationKmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-automaticbackupreplicationkmskeyid
	AutomaticBackupReplicationKmsKeyId *types.Value `json:"AutomaticBackupReplicationKmsKeyId,omitempty"`

	// AutomaticBackupReplicationRegion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-automaticbackupreplicationregion
	AutomaticBackupReplicationRegion *types.Value `json:"AutomaticBackupReplicationRegion,omitempty"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-availabilityzone
	AvailabilityZone *types.Value `json:"AvailabilityZone,omitempty"`

	// BackupRetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-backupretentionperiod
	BackupRetentionPeriod *types.Value `json:"BackupRetentionPeriod,omitempty"`

	// CACertificateIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-cacertificateidentifier
	CACertificateIdentifier *types.Value `json:"CACertificateIdentifier,omitempty"`

	// CertificateDetails AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-certificatedetails
	CertificateDetails *DBInstance_CertificateDetails `json:"CertificateDetails,omitempty"`

	// CertificateRotationRestart AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-certificaterotationrestart
	CertificateRotationRestart *types.Value `json:"CertificateRotationRestart,omitempty"`

	// CharacterSetName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-charactersetname
	CharacterSetName *types.Value `json:"CharacterSetName,omitempty"`

	// CopyTagsToSnapshot AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-copytagstosnapshot
	CopyTagsToSnapshot *types.Value `json:"CopyTagsToSnapshot,omitempty"`

	// CustomIAMInstanceProfile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-customiaminstanceprofile
	CustomIAMInstanceProfile *types.Value `json:"CustomIAMInstanceProfile,omitempty"`

	// DBClusterIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbclusteridentifier
	DBClusterIdentifier *types.Value `json:"DBClusterIdentifier,omitempty"`

	// DBClusterSnapshotIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbclustersnapshotidentifier
	DBClusterSnapshotIdentifier *types.Value `json:"DBClusterSnapshotIdentifier,omitempty"`

	// DBInstanceClass AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbinstanceclass
	DBInstanceClass *types.Value `json:"DBInstanceClass,omitempty"`

	// DBInstanceIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbinstanceidentifier
	DBInstanceIdentifier *types.Value `json:"DBInstanceIdentifier,omitempty"`

	// DBName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbname
	DBName *types.Value `json:"DBName,omitempty"`

	// DBParameterGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbparametergroupname
	DBParameterGroupName *types.Value `json:"DBParameterGroupName,omitempty"`

	// DBSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbsecuritygroups
	DBSecurityGroups *types.Value `json:"DBSecurityGroups,omitempty"`

	// DBSnapshotIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbsnapshotidentifier
	DBSnapshotIdentifier *types.Value `json:"DBSnapshotIdentifier,omitempty"`

	// DBSubnetGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbsubnetgroupname
	DBSubnetGroupName *types.Value `json:"DBSubnetGroupName,omitempty"`

	// DBSystemId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dbsystemid
	DBSystemId *types.Value `json:"DBSystemId,omitempty"`

	// DedicatedLogVolume AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-dedicatedlogvolume
	DedicatedLogVolume *types.Value `json:"DedicatedLogVolume,omitempty"`

	// DeleteAutomatedBackups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-deleteautomatedbackups
	DeleteAutomatedBackups *types.Value `json:"DeleteAutomatedBackups,omitempty"`

	// DeletionProtection AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-deletionprotection
	DeletionProtection *types.Value `json:"DeletionProtection,omitempty"`

	// Domain AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-domain
	Domain *types.Value `json:"Domain,omitempty"`

	// DomainAuthSecretArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-domainauthsecretarn
	DomainAuthSecretArn *types.Value `json:"DomainAuthSecretArn,omitempty"`

	// DomainDnsIps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-domaindnsips
	DomainDnsIps *types.Value `json:"DomainDnsIps,omitempty"`

	// DomainFqdn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-domainfqdn
	DomainFqdn *types.Value `json:"DomainFqdn,omitempty"`

	// DomainIAMRoleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-domainiamrolename
	DomainIAMRoleName *types.Value `json:"DomainIAMRoleName,omitempty"`

	// DomainOu AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-domainou
	DomainOu *types.Value `json:"DomainOu,omitempty"`

	// EnableCloudwatchLogsExports AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-enablecloudwatchlogsexports
	EnableCloudwatchLogsExports *types.Value `json:"EnableCloudwatchLogsExports,omitempty"`

	// EnableIAMDatabaseAuthentication AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-enableiamdatabaseauthentication
	EnableIAMDatabaseAuthentication *types.Value `json:"EnableIAMDatabaseAuthentication,omitempty"`

	// EnablePerformanceInsights AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-enableperformanceinsights
	EnablePerformanceInsights *types.Value `json:"EnablePerformanceInsights,omitempty"`

	// Endpoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-endpoint
	Endpoint *DBInstance_Endpoint `json:"Endpoint,omitempty"`

	// Engine AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-engine
	Engine *types.Value `json:"Engine,omitempty"`

	// EngineLifecycleSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-enginelifecyclesupport
	EngineLifecycleSupport *types.Value `json:"EngineLifecycleSupport,omitempty"`

	// EngineVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-engineversion
	EngineVersion *types.Value `json:"EngineVersion,omitempty"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-iops
	Iops *types.Value `json:"Iops,omitempty"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-kmskeyid
	KmsKeyId *types.Value `json:"KmsKeyId,omitempty"`

	// LicenseModel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-licensemodel
	LicenseModel *types.Value `json:"LicenseModel,omitempty"`

	// ManageMasterUserPassword AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-managemasteruserpassword
	ManageMasterUserPassword *types.Value `json:"ManageMasterUserPassword,omitempty"`

	// MasterUserPassword AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-masteruserpassword
	MasterUserPassword *types.Value `json:"MasterUserPassword,omitempty"`

	// MasterUserSecret AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-masterusersecret
	MasterUserSecret *DBInstance_MasterUserSecret `json:"MasterUserSecret,omitempty"`

	// MasterUsername AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-masterusername
	MasterUsername *types.Value `json:"MasterUsername,omitempty"`

	// MaxAllocatedStorage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-maxallocatedstorage
	MaxAllocatedStorage *types.Value `json:"MaxAllocatedStorage,omitempty"`

	// MonitoringInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-monitoringinterval
	MonitoringInterval *types.Value `json:"MonitoringInterval,omitempty"`

	// MonitoringRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-monitoringrolearn
	MonitoringRoleArn *types.Value `json:"MonitoringRoleArn,omitempty"`

	// MultiAZ AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-multiaz
	MultiAZ *types.Value `json:"MultiAZ,omitempty"`

	// NcharCharacterSetName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-ncharcharactersetname
	NcharCharacterSetName *types.Value `json:"NcharCharacterSetName,omitempty"`

	// NetworkType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-networktype
	NetworkType *types.Value `json:"NetworkType,omitempty"`

	// OptionGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-optiongroupname
	OptionGroupName *types.Value `json:"OptionGroupName,omitempty"`

	// PerformanceInsightsKMSKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-performanceinsightskmskeyid
	PerformanceInsightsKMSKeyId *types.Value `json:"PerformanceInsightsKMSKeyId,omitempty"`

	// PerformanceInsightsRetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-performanceinsightsretentionperiod
	PerformanceInsightsRetentionPeriod *types.Value `json:"PerformanceInsightsRetentionPeriod,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-port
	Port *types.Value `json:"Port,omitempty"`

	// PreferredBackupWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-preferredbackupwindow
	PreferredBackupWindow *types.Value `json:"PreferredBackupWindow,omitempty"`

	// PreferredMaintenanceWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-preferredmaintenancewindow
	PreferredMaintenanceWindow *types.Value `json:"PreferredMaintenanceWindow,omitempty"`

	// ProcessorFeatures AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-processorfeatures
	ProcessorFeatures []DBInstance_ProcessorFeature `json:"ProcessorFeatures,omitempty"`

	// PromotionTier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-promotiontier
	PromotionTier *types.Value `json:"PromotionTier,omitempty"`

	// PubliclyAccessible AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-publiclyaccessible
	PubliclyAccessible *types.Value `json:"PubliclyAccessible,omitempty"`

	// ReplicaMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-replicamode
	ReplicaMode *types.Value `json:"ReplicaMode,omitempty"`

	// RestoreTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-restoretime
	RestoreTime *types.Value `json:"RestoreTime,omitempty"`

	// SourceDBClusterIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-sourcedbclusteridentifier
	SourceDBClusterIdentifier *types.Value `json:"SourceDBClusterIdentifier,omitempty"`

	// SourceDBInstanceAutomatedBackupsArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-sourcedbinstanceautomatedbackupsarn
	SourceDBInstanceAutomatedBackupsArn *types.Value `json:"SourceDBInstanceAutomatedBackupsArn,omitempty"`

	// SourceDBInstanceIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-sourcedbinstanceidentifier
	SourceDBInstanceIdentifier *types.Value `json:"SourceDBInstanceIdentifier,omitempty"`

	// SourceDbiResourceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-sourcedbiresourceid
	SourceDbiResourceId *types.Value `json:"SourceDbiResourceId,omitempty"`

	// SourceRegion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-sourceregion
	SourceRegion *types.Value `json:"SourceRegion,omitempty"`

	// StorageEncrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-storageencrypted
	StorageEncrypted *types.Value `json:"StorageEncrypted,omitempty"`

	// StorageThroughput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-storagethroughput
	StorageThroughput *types.Value `json:"StorageThroughput,omitempty"`

	// StorageType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-storagetype
	StorageType *types.Value `json:"StorageType,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-tags
	Tags []cloudformation.Tag `json:"Tags,omitempty"`

	// Timezone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-timezone
	Timezone *types.Value `json:"Timezone,omitempty"`

	// UseDefaultProcessorFeatures AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-usedefaultprocessorfeatures
	UseDefaultProcessorFeatures *types.Value `json:"UseDefaultProcessorFeatures,omitempty"`

	// UseLatestRestorableTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-uselatestrestorabletime
	UseLatestRestorableTime *types.Value `json:"UseLatestRestorableTime,omitempty"`

	// VPCSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-vpcsecuritygroups
	VPCSecurityGroups *types.Value `json:"VPCSecurityGroups,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DBInstance) AWSCloudFormationType() string {
	return "AWS::RDS::DBInstance"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r DBInstance) MarshalJSON() ([]byte, error) {
	type Properties DBInstance
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *DBInstance) UnmarshalJSON(b []byte) error {
	type Properties DBInstance
	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           []string
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = DBInstance(*res.Properties)
	}
	if res.DependsOn != nil {
		r.AWSCloudFormationDependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
