FROM golang:1.10.1-alpine
# ARG version
WORKDIR /go/src/github.com/nicksnyder/service
COPY . .
RUN go install \
    -v \
    # -ldflags "-X github.com/nicksnyder/service/pkg/version.version=$version" \
    ./cmd/config \
    ./cmd/frontend
RUN ls -lh /go/bin