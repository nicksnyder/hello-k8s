# This dockerfile builds all binaries for the project.
FROM golang:1.13.4-alpine
WORKDIR /go/src/github.com/nicksnyder/hello-server
COPY . .
ARG version
RUN test -n "$version"
RUN go install \
    -v \
    -ldflags "-X github.com/nicksnyder/hello-server/cmd/internal/binary.version=$version" \
    ./cmd/config \
    ./cmd/frontend