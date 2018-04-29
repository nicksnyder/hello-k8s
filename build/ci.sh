#!/bin/sh

set -e

cd "$(dirname "${BASH_SOURCE[0]}")/.."

go test ./...

DOCKER_TAG=`./build/docker_tag.sh`
time docker build \
    -f ./build/base.Dockerfile \
    -t base:$DOCKER_TAG \
    -t base:latest \
    --build-arg dockerTag=$DOCKER_TAG \
    .

time docker build \
    -f ./build/frontend.Dockerfile \
    -t nickdsnyder/frontend:$DOCKER_TAG \
    -t nickdsnyder/frontend:latest \
    .

time docker build \
    -f ./build/config.Dockerfile \
    -t nickdsnyder/config:$DOCKER_TAG \
    -t nickdsnyder/config:latest \
    .

echo "Built $DOCKER_TAG"

# if test -n "$(git status --porcelain)"; then
#     echo "Skipping push due to dirty working copy"
#     exit
# fi

time docker push nickdsnyder/frontend:$DOCKER_TAG
time docker push nickdsnyder/config:$DOCKER_TAG
time docker push nickdsnyder/frontend:latest
time docker push nickdsnyder/config:latest

echo "Pushed $DOCKER_TAG"

./deploy/apply.sh \
    --set frontend.version=$DOCKER_TAG \
    --set config.version=$DOCKER_TAG

echo "Deployed $DOCKER_TAG"