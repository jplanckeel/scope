FROM golang:alpine as builder

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

# Update package
RUN apk update && \
    apk upgrade

# hadolint ignore=DL3018
RUN apk add --no-cache curl

COPY --from=builder /go/scope /bin/scope

CMD ["/bin/sh"]
