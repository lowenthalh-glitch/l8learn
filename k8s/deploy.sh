#!/usr/bin/env bash
set -e
echo "Deploying L8Learn..."
kubectl apply -f ./vnet.yaml
kubectl apply -f ./learn.yaml
kubectl apply -f ./web.yaml
echo "L8Learn deployed."
