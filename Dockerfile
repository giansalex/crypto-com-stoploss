FROM golang:1.14-alpine as builder

WORKDIR /root

RUN apk update && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

# Copy the source code to workspace.
COPY . .

# Build executable
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -gcflags=all=-l ./app -o /build/crypto

# Release image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build .

ENTRYPOINT ["./crypto"]