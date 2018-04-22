#!/bin/sh

cd "$(dirname "${BASH_SOURCE[0]}")/.."

time docker build -f ./build/base.Dockerfile -t base .
time docker build -f ./build/frontend.Dockerfile -t nickdsnyder/frontend .
time docker build -f ./build/config.Dockerfile -t nickdsnyder/config .