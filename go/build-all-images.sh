#!/usr/bin/env bash
set -e
echo "Building all L8Learn images..."

echo "--- Building learn-vnet ---"
cd learn/vnet && ./build.sh && cd ../..

echo "--- Building learn-logs-vnet ---"
cd learn/log-vnet && ./build.sh && cd ../..

echo "--- Building learn (backend) ---"
cd learn/main && ./build.sh && cd ../..

echo "--- Building learn-web (UI) ---"
cd learn/ui && ./build.sh && cd ../..

echo "--- Building learn-log-agent ---"
cd learn/log-agent && ./build.sh && cd ../..

echo "All images built and pushed."
