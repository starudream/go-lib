#!/usr/bin/env bash

FILES=(
    ## https://github.com/protocolbuffers/protobuf
    protocolbuffers/protobuf/raw/main/src@google/protobuf/descriptor.proto

    protocolbuffers/protobuf/raw/main/src@google/protobuf/any.proto
    protocolbuffers/protobuf/raw/main/src@google/protobuf/duration.proto
    protocolbuffers/protobuf/raw/main/src@google/protobuf/empty.proto
    protocolbuffers/protobuf/raw/main/src@google/protobuf/struct.proto
    protocolbuffers/protobuf/raw/main/src@google/protobuf/timestamp.proto
    protocolbuffers/protobuf/raw/main/src@google/protobuf/wrappers.proto

    ## https://github.com/googleapis/googleapis
    googleapis/googleapis/raw/master@google/api/annotations.proto
    googleapis/googleapis/raw/master@google/api/http.proto

    ## https://github.com/grpc-ecosystem/grpc-gateway
    grpc-ecosystem/grpc-gateway/raw/main@protoc-gen-openapiv2/options/annotations.proto
    grpc-ecosystem/grpc-gateway/raw/main@protoc-gen-openapiv2/options/openapiv2.proto

    ## https://github.com/bufbuild/protoc-gen-validate
    bufbuild/protoc-gen-validate/raw/main@validate/validate.proto
)

for file in "${FILES[@]}" ; do
    a=$(echo "${file}" | cut -d "@" -f1)
    b=$(echo "${file}" | cut -d "@" -f2)
    echo "Downloading ${b}"
    mkdir -p "$(dirname "${b}")"
    wget -qO "${b}" "https://github.com/${a}/${b}"
done
