#!/usr/bin/env bash
#
# L8Learn — Local Development Setup
# Adapted from ../l8erp/go/run-local.sh
#
set -e

echo "=== L8Learn Local Development ==="

# Clean and fetch dependencies
echo "Cleaning dependencies..."
rm -rf go.sum go.mod vendor
go mod init
GOPROXY=direct GOPRIVATE=github.com go mod tidy
go mod vendor

# Start database
echo "Starting database..."
docker rm -f unsecure-postgres 2>/dev/null || true
docker run -d --name unsecure-postgres -p 5432:5432 -v /data/:/data/ \
    saichler/unsecure-postgres:latest \
    /bin/sh -c "/start-postgres.sh admin admin admin 5432 && tail -f /dev/null"
sleep 5

# Build binaries
echo "Building binaries..."
rm -rf demo && mkdir -p demo

echo "  Building mocks..."
cd tests/mocks/cmd && go build -o ../../../demo/mocks_demo && cd ../../..

echo "  Building vnet..."
cd learn/vnet && go build -o ../../demo/vnet_demo && cd ../..

echo "  Building backend..."
cd learn/main && go build -o ../../demo/learn_demo && cd ../..

echo "  Building UI server..."
cd learn/ui/main && go build -o ../../../demo/ui_demo && cd ../../..

# Copy web assets
echo "  Copying web assets..."
cp -r learn/ui/web demo/web

# Generate cleanup script
cat > demo/kill_demo.sh << 'KILLEOF'
#!/usr/bin/env bash
pkill -f "vnet_demo" 2>/dev/null || true
pkill -f "learn_demo" 2>/dev/null || true
pkill -f "ui_demo" 2>/dev/null || true
docker rm -f unsecure-postgres 2>/dev/null || true
echo "Demo processes killed."
KILLEOF
chmod +x demo/kill_demo.sh

# Start services
echo ""
echo "Starting services..."
cd demo

./vnet_demo &
sleep 1

./learn_demo local &
./ui_demo &
sleep 8

# Get external IP
EXTERNAL_IP=$(hostname -I | awk '{print $1}')
echo ""
echo "================================================"
echo "  L8Learn is running!"
echo "  URL: https://${EXTERNAL_IP}:2773"
echo "  Login: admin / admin"
echo "================================================"
echo ""

# Upload mock data
read -p "Press ENTER to upload mock data (or Ctrl+C to skip)..."
./mocks_demo --address https://${EXTERNAL_IP}:2773 --user admin --password admin --insecure

echo ""
read -p "Press ENTER to shut down..."
./kill_demo.sh
