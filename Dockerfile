FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/jplanckeel/scope

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Build the package
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /go/scope


FROM golang:alpine as prod

ENV HELM_VERSION=3.5.4

# Update package
RUN apk update && \
    apk upgrade

# hadolint ignore=DL3018
RUN apk add --no-cache curl

COPY --from=builder /go/scope /bin/scope

RUN /bin/sh -c cd /tmp/ && \
    wget https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz && \
    tar -zxf helm-v*-linux-amd64.tar.gz && mv linux-amd64/helm /usr/local/bin/helm && \
    rm -rf /tmp/*.tar.gz /tmp/linux-amd64

RUN adduser -D scope
USER scope

ENTRYPOINT  [ "/bin/scope" ]
