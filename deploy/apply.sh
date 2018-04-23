#!/bin/sh

cd "$(dirname "${BASH_SOURCE[0]}")"

helm template ./chart | kubectl apply -f -