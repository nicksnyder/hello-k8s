#!/bin/sh

./build.sh
time docker push nickdsnyder/frontend
time docker push nickdsnyder/config