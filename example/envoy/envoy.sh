#!/usr/bin/env bash

docker run -d \
    --name envoy \
    --restart always \
    -p "9901:9901" \
    -p "10000:10000" \
    -v "$(pwd)"/envoy.yaml:/opt/bitnami/envoy/conf/envoy.yaml \
    --add-host host.docker.internal:host-gateway \
    bitnami/envoy:1.29.2
