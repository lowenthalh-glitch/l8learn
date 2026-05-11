#!/usr/bin/env bash

set -e

wget https://raw.githubusercontent.com/saichler/l8types/refs/heads/main/proto/api.proto
wget https://raw.githubusercontent.com/saichler/l8common/refs/heads/main/proto/l8common.proto

# Generate bindings for all learn proto files
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-content.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-students.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-adaptive.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-assessment.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-analytics.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-history.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-homeschool.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-collab.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-profile.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-llm.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-eval.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=learn-generated.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest

rm api.proto l8common.proto

# Move generated bindings to go/types and clean up
rm -rf ../go/types
mkdir -p ../go/types
rm -rf ./types/l8common
mv ./types/* ../go/types/.
rm -rf ./types

rm -rf *.rs

# Fix import paths
cd ../go
find . -name "*.go" -type f -exec sed -i 's|"./types/l8services"|"github.com/saichler/l8types/go/types/l8services"|g' {} +
find . -name "*.go" -type f -exec sed -i 's|"./types/l8api"|"github.com/saichler/l8types/go/types/l8api"|g' {} +
find . -name "*.go" -type f -exec sed -i 's|"./types/l8common"|"github.com/saichler/l8common/go/types/l8common"|g' {} +
find . -name "*.go" -type f -exec sed -i 's|"./types/learn"|"github.com/saichler/l8learn/go/types/learn"|g' {} +
