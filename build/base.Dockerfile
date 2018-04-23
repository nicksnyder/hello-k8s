# This dockerfile builds all binaries for the project.
FROM golang:1.10.1-alpine
# ARG version
WORKDIR /go/src/github.com/nicksnyder/hello-server
COPY . .
RUN go install \
    -v \
    # -ldflags "-X github.com/nicksnyder/hello-server/pkg/version.version=$version" \
    ./cmd/config \
    ./cmd/frontend
RUN ls -lh /go/bin