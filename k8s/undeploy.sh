#!/usr/bin/env bash
set -e
echo "Undeploying L8Learn..."
kubectl delete -f ./web.yaml 2>/dev/null || true
kubectl delete -f ./learn.yaml 2>/dev/null || true
kubectl delete -f ./vnet.yaml 2>/dev/null || true
echo "L8Learn undeployed."
