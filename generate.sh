#!/usr/bin/env bash

protoc --proto_path=proto_files proto_files/*.proto --go_out=paths=source_relative:./plugnmeet