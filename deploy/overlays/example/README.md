This example customizes the number of frontend replicas.

You can deploy with `kustomize build deploy/overlays/example/ | kubectl apply --prune -l app=hello-server -f -`.
