#!/usr/bin/env bash
set -e
echo "Undeploying L8Learn..."
kubectl delete -f ./learn-local.yaml 2>/dev/null || true
echo "L8Learn undeployed."
