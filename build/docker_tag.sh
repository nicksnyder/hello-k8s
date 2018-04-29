#!/bin/sh

DIRTY=`test -z "$(git status --porcelain)" || echo "-dirty"`
echo `date -u +"%Y-%m-%d"`-`git rev-parse --short HEAD`$DIRTY