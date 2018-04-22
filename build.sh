#!/bin/sh

time docker build -f base.Dockerfile -t base .
time docker build -f frontend.Dockerfile -t frontend .
time docker build -f config.Dockerfile -t config .
