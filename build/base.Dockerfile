# This dockerfile builds all binaries for the project.
FROM golang:1.10.1-alpine
WORKDIR /go/src/github.com/nicksnyder/hello-server
COPY . .
ARG dockerTag
RUN test -n "$dockerTag"
RUN go install \
    -v \
    -ldflags "-X github.com/nicksnyder/hello-server/pkg/binary.dockerTag=$dockerTag" \
    ./cmd/config \
    ./cmd/frontend