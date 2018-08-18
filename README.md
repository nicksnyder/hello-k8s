# Hello server

This is a test application with multiple services so I can experiment with Kubernetes and Docker.

## Build

Each commit builds and pushes a tagged docker image for each service.

Each commit to master pushes a `:insiders` docker image for each service.

Each semver tag pushes a `:vX.Y.Z` docker image for each service.

## Deploy

This service can be deployed to Kubernetes with `./deploy/apply.sh`.
