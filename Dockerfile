<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
ARG BUILD_IMAGE=public.ecr.aws/eksctl/eksctl-build:833f4464e865a6398788bf6cbc5447967b8974b7
=======
ARG BUILD_IMAGE=public.ecr.aws/eksctl/eksctl-build:5e2d110f32a60d699ca0642ac3287190227c45ec
>>>>>>> 3370879ad (Update build image manifest, tag file and workflows)
=======
ARG BUILD_IMAGE=public.ecr.aws/eksctl/eksctl-build:9b3f575ceb1a272a000ca1fbaded0174096a007a
>>>>>>> ec7882c9c (Update build image go version to 1.21)
=======
ARG BUILD_IMAGE=public.ecr.aws/eksctl/eksctl-build:79ff6e4d9b5ba7e2bfb962efe2fcb7b2eebc1f2a
>>>>>>> ac7fabff7 (Fix rebase)
FROM $BUILD_IMAGE as build

WORKDIR /src

COPY . /src

RUN make test
RUN make build \
    && cp ./eksctl /out/usr/local/bin/eksctl
RUN make build-integration-test \
    && mkdir -p /out/usr/local/share/eksctl \
    && cp -r integration/data/*.yaml integration/scripts /out/usr/local/share/eksctl \
    && cp ./eksctl-integration-test /out/usr/local/bin/eksctl-integration-test

FROM scratch
COPY --from=build /out /
ENTRYPOINT ["eksctl"]
