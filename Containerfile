FROM docker.io/library/golang:1.16 as builder
WORKDIR /tmd
COPY . .
RUN go mod tidy; \
    go build
FROM registry.fedoraproject.org/fedora:35
WORKDIR /tmd
COPY --from=builder /tmd/local-api-global-variables ./local-api-global-variables
RUN chmod +x local-api-global-variables
ENTRYPOINT ["./local-api-global-variables"]
