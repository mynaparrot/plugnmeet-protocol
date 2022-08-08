#!/usr/bin/env bash

protoc \
-I ${GOPATH}/src \
-I ${GOPATH}/pkg/mod \
--proto_path=proto_files \
proto_files/*.proto \
--go_out=paths=source_relative:./plugnmeet \
--validate_out="lang=go,paths=source_relative:./plugnmeet"
