#!/usr/bin/env bash
#
# Copyright 2026 Sharon Aicler (saichler@gmail.com)
# Licensed under the Apache License, Version 2.0
#
# Generates all protobuf bindings for l8learn.
# Run from the proto/ directory: cd proto && ./make-bindings.sh
#
set -e

echo "Downloading api.proto..."
wget -q -O api.proto https://raw.githubusercontent.com/saichler/l8types/master/proto/api.proto 2>/dev/null || true

echo "Downloading l8common.proto..."
wget -q -O l8common.proto https://raw.githubusercontent.com/saichler/l8common/master/proto/l8common.proto 2>/dev/null || true

for proto in learn-content learn-students learn-adaptive learn-assessment learn-analytics learn-history learn-homeschool learn-collab; do
    echo "Compiling ${proto}.proto..."
    docker run -i --rm -v $(pwd):/proto -w /proto saichler/protoc:latest \
        protoc --go_out=. --go_opt=paths=source_relative \
        ${proto}.proto
done

echo "Moving generated files to ../go/types/learn/..."
mkdir -p ../go/types/learn
mv -f types/learn/*.pb.go ../go/types/learn/ 2>/dev/null || true
rm -rf types/

echo "Done! Generated .pb.go files in go/types/learn/"
