# Hello server

This is a multi-service server application. There are two services:

- Frontend
- Config

## Build

Each commit builds and pushes a tagged docker image for each service.

Each commit to master pushes a `:insiders` docker image for each service.

Each semver tag pushes a `:vX.Y.Z` docker image for each service.

## Deploy

This service can be deployed to Kubernetes with `./deploy/apply.sh`.

Use [Kustomize](https://github.com/kubernetes-sigs/kustomize) if you need to customize the yaml.
