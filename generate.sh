#!/usr/bin/env sh

PROTO_TMP_DIR="./tmp"
PROTO_VALIDATOR_DIR="${PROTO_TMP_DIR}/protovalidate"
LIVEKIT_PROTO_DIR="${PROTO_TMP_DIR}/livekit-protocol"

if [ ! -d ${PROTO_TMP_DIR} ]; then
  mkdir -p ${PROTO_TMP_DIR}
fi

if [ ! -d ${PROTO_VALIDATOR_DIR} ]; then
  git clone https://github.com/bufbuild/protovalidate ${PROTO_VALIDATOR_DIR}
fi

if [ ! -d ${LIVEKIT_PROTO_DIR} ]; then
  git clone https://github.com/livekit/protocol ${LIVEKIT_PROTO_DIR}
fi


if ! which buf >/dev/null; then
  printf "buf not installed, installing using go install"
  go install github.com/bufbuild/buf/cmd/buf@latest
fi

buf dep update
buf generate

cd js
./generate.sh
