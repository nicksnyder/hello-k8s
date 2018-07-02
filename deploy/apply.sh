#!/bin/sh

kubectl apply --prune -l app=hello-server -f ./deploy/base/