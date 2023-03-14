#!/usr/bin/env bash
GOPATH=$HOME/go

protoc \
-I ${GOPATH}/pkg/mod \
-I ${GOPATH}/pkg/mod/github.com/livekit/protocol \
-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate \
--proto_path=proto_files \
proto_files/*.proto \
--go_out=paths=source_relative:./plugnmeet \
--validate_out="lang=go,paths=source_relative:./plugnmeet"
