#!/usr/bin/env bash
set -e
docker build --no-cache --platform=linux/amd64 -t saichler/learn-web:latest .
docker push saichler/learn-web:latest
