#!/bin/bash
# Echos a version tag for this build.
# If the current commit has a semver tag pointing to it, use it.
# Otherwise, generate a tag that contains the date and commit sha.

cd "$(dirname "${BASH_SOURCE[0]}")/.." # cd to repo root dir

# Add -dirty suffix to version if there are uncommitted changes in the working copy.
dirty=`test -z "$(git status --porcelain)" || echo "-dirty"`

tagged_version=`git tag --points-at HEAD | tail -n 1`
generated_version=`date -u +"%Y-%m-%d"`-`git rev-parse --short HEAD`

./build/semver.sh $tagged_version && echo $tagged_version$dirty || echo $generated_version$dirty