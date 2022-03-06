FROM docker.io/library/golang:1.16 as builder
WORKDIR /tmd
RUN go mod vendor; \
    go build

FROM registry.fedoraproject.org/fedora:35
WORKDIR /tmd
ENTRYPOINT ["./local-api-global-variables"]
