FROM docker.io/library/golang:1.16 as builder
WORKDIR /tmd
COPY . .
RUN go mod vendor; \
    go build

FROM registry.fedoraproject.org/fedora:35
WORKDIR /tmd
COPY --from=builder --chmod=700 --chown=0:0 /tmd/local-api-global-variables .
ENTRYPOINT ["./local-api-global-variables"]
