#!/usr/bin/env bash

protoc \
-I ${GOPATH}/src \
-I ${GOPATH}/pkg/mod \
-I ${GOPATH}/pkg/mod/github.com/livekit/protocol@v1.3.1 \
-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.9.1 \
--proto_path=proto_files \
proto_files/*.proto \
--go_out=paths=source_relative:./plugnmeet \
--validate_out="lang=go,paths=source_relative:./plugnmeet"
