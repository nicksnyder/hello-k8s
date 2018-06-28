# This dockerfile builds all binaries for the project.
FROM golang:1.10.3-alpine
WORKDIR /go/src/github.com/nicksnyder/hello-server
COPY . .
ARG version
RUN test -n "$version"
RUN go install \
    -v \
    -ldflags "-X github.com/nicksnyder/hello-server/pkg/binary.version=$version" \
    ./cmd/config \
    ./cmd/frontend