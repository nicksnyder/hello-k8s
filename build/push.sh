#!/bin/sh

cd "$(dirname "${BASH_SOURCE[0]}")/.."

./build/build.sh
time docker push nickdsnyder/frontend
time docker push nickdsnyder/config
