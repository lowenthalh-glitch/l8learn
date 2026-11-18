#!/usr/bin/env bash
set -e
cp ~/.netrc .netrc
docker build --no-cache --platform=linux/amd64 -t registry.layer-8.dev:44443/l8learn/l8learn-log-agent:latest .
rm -f .netrc
docker push registry.layer-8.dev:44443/l8learn/l8learn-log-agent:latest
