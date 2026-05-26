#!/usr/bin/env bash
set -e
cp ~/.netrc .netrc
docker build --no-cache --platform=linux/amd64 -t saichler/l8learn-web:latest .
rm -f .netrc
docker push saichler/l8learn-web:latest
