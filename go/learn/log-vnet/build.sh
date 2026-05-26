#!/usr/bin/env bash
set -e
docker build --no-cache --platform=linux/amd64 -t saichler/l8learn-logs-vnet:latest .
docker push saichler/l8learn-logs-vnet:latest
