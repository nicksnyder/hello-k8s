#!/bin/sh
# Exits with non-zero status if the first parameter does not look like a semver.
# Regex is permissive with respect to prerelease suffix, I am not trying to win the regex olympics.
semver='^v[0-9]+(\.[0-9]+){0,2}(-[0-9A-Za-z.-]+)?$'
[[ $1 =~ $semver ]] || exit 1