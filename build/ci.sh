#!/bin/bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")/.." # cd to repo root dir

go test ./...

version=`./build/version.sh`
time docker build \
    -f ./build/base.Dockerfile \
    -t base:latest \
    --build-arg version=$version \
    .

time docker build \
    -f ./build/frontend.Dockerfile \
    -t nickdsnyder/frontend:$version \
    -t nickdsnyder/frontend:insiders \
    .

time docker build \
    -f ./build/config.Dockerfile \
    -t nickdsnyder/config:$version \
    -t nickdsnyder/config:insiders \
    .

echo "Built $version"

# if test -n "$(git status --porcelain)"; then
#     echo "Skipping push due to dirty working copy"
#     exit
# fi

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

if [ "$TRAVIS_BRANCH" = "master" ] ; then
    time docker push nickdsnyder/frontend:$version
    time docker push nickdsnyder/config:$version
    echo "Pushed :$version"

    time docker push nickdsnyder/frontend:insiders
    time docker push nickdsnyder/config:insiders
    echo "Pushed :insiders"
fi

if ./build/semver.sh $version; then
    time docker push nickdsnyder/frontend:latest
    time docker push nickdsnyder/config:latest
    echo "Pushed :latest"
fi


# This is how auto-deploy would work from CI.
# ./deploy/apply.sh \
#     --set frontend.version=$version \
#     --set config.version=$version

# echo "Deployed $version"