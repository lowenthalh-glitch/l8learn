#!/bin/bash
set -e

# Always kill previous demo processes
pkill -9 demo 2>/dev/null || true
sleep 1

# Vendor refresh
rm -rf go.sum go.mod vendor
go mod init
GOPROXY=direct GOPRIVATE=github.com go mod tidy
go mod vendor

# Recreate database
echo "Starting database..."
docker rm -f unsecure-postgres 2>/dev/null || true
docker run -d --name unsecure-postgres -p 5432:5432 \
    -v /data/:/data/ saichler/unsecure-postgres:latest \
    /bin/sh -c "/start-postgres.sh admin admin admin 5432 && tail -f /dev/null"
sleep 5

# Build binaries
rm -rf demo && mkdir -p demo
cd tests/mocks/cmd && go build -o ../../../demo/mocks_demo && cd ../../../
cd learn/vnet && go build -o ../../demo/vnet_demo && cd ../../
cd learn/main && go build -o ../../demo/learn_demo && cd ../../
cd learn/ui/main && go build -o ../../../demo/ui_demo && cd ../../../
cp -r learn/ui/web demo/.

# Generate kill script
cd demo
cat > kill_demo.sh <<'EOF'
cd ..
rm -rf demo
rm -rf /data/postgres/admin
pkill -9 demo
EOF
chmod +x kill_demo.sh

LOGFILE="../demo/demo.log"
> $LOGFILE

# Start services — UI must start BEFORE backend so it receives web service broadcasts
echo "Starting VNet..."
./vnet_demo >> $LOGFILE 2>&1 &
sleep 2

echo "Starting UI..."
./ui_demo >> $LOGFILE 2>&1 &
sleep 2

echo "Starting backend (local mode)..."
./learn_demo local >> $LOGFILE 2>&1 &
sleep 5

echo "All services started! Logs: demo/demo.log"

# Upload mock data
EXTERNAL_IP=$(ip route get 1.1.1.1 | grep -oP 'src \K[0-9.]+')
echo "Mock data modes: demo (default), security, full"
read -p "Enter mode [demo]: " MOCK_MODE
MOCK_MODE=${MOCK_MODE:-demo}
./mocks_demo --address https://${EXTERNAL_IP}:2773 --user admin --password admin --insecure --mode ${MOCK_MODE} 2>&1 | tee -a $LOGFILE

read -p "Press Enter to kill the demo"
./kill_demo.sh
