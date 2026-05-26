#!/usr/bin/env bash
set -e
echo "Deploying L8Learn..."
kubectl apply -f ./learn-local.yaml
echo "L8Learn deployed."
