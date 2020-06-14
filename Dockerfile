FROM golang:alpine3.11 as builder

WORKDIR /go/src/github.com/giansalex/crypto-com-trailing-stop-loss

RUN apk update && apk add curl git

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/giansalex/crypto-com-trailing-stop-loss

# Install packages
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure

# Build executable
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -gcflags=all=-l -o /build/crypto

# Release image
FROM scratch

COPY --from=builder /build .

ENTRYPOINT ["./crypto"]