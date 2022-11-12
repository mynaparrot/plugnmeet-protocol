#!/usr/bin/env bash

protoc \
-I ${GOPATH}/src \
-I ${GOPATH}/pkg/mod \
-I ${GOPATH}/pkg/mod/github.com/livekit/protocol@v1.2.2 \
-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.9.0 \
--proto_path=proto_files \
proto_files/*.proto \
--go_out=paths=source_relative:./plugnmeet \
--validate_out="lang=go,paths=source_relative:./plugnmeet"
