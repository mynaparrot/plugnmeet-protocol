#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# Absolute tmp dir; already gitignored.
PROTO_TMP_DIR="${SCRIPT_DIR}/tmp"

# -------------------------------------------------------------------
# Ensure required tooling is available
# -------------------------------------------------------------------
if ! which jq >/dev/null 2>&1; then
  echo "ERROR: jq is required but not installed."
  echo "       Install it (e.g. 'apt install jq' or 'brew install jq') and re-run."
  exit 1
fi

if ! which go >/dev/null 2>&1; then
  echo "ERROR: go is required but not installed."
  echo "       Install it (e.g. from https://go.dev/dl/) and re-run."
  exit 1
fi

if ! which buf >/dev/null 2>&1; then
  echo "buf not found, installing using go install"
  go install github.com/bufbuild/buf/cmd/buf@latest
fi

# -------------------------------------------------------------------
# Clean and recreate tmp
# -------------------------------------------------------------------
rm -rf "${PROTO_TMP_DIR}"
mkdir -p "${PROTO_TMP_DIR}"

# -------------------------------------------------------------------
# Resolve livekit/protocol version from go.mod → Go module cache
# -------------------------------------------------------------------
LIVEKIT_VERSION=$(grep 'github.com/livekit/protocol ' go.mod | awk '{print $2}')
echo "→ Resolving livekit/protocol ${LIVEKIT_VERSION} from Go module cache..."
go mod download
LIVEKIT_DIR=$(go list -m -json "github.com/livekit/protocol@${LIVEKIT_VERSION}" | jq -r .Dir)
echo "  ${LIVEKIT_DIR}"

if [ ! -d "${LIVEKIT_DIR}/protobufs" ]; then
  echo "ERROR: livekit protobufs not found at ${LIVEKIT_DIR}/protobufs"
  exit 1
fi

# -------------------------------------------------------------------
# Export protovalidate protos from BSR (latest commit)
# -------------------------------------------------------------------
PROTOVALIDATE_DIR="${PROTO_TMP_DIR}/protovalidate-proto"
echo "→ Exporting buf.build/bufbuild/protovalidate from BSR..."
buf export "buf.build/bufbuild/protovalidate" -o "${PROTOVALIDATE_DIR}"
echo "  ${PROTOVALIDATE_DIR}"

# -------------------------------------------------------------------
# Create symlinks so buf.yaml relative paths resolve correctly
# -------------------------------------------------------------------
# livekit: buf.yaml expects ./tmp/livekit-protocol/protobufs
mkdir -p "${PROTO_TMP_DIR}/livekit-protocol"
ln -sf "${LIVEKIT_DIR}/protobufs" "${PROTO_TMP_DIR}/livekit-protocol/protobufs"

# protovalidate: buf.yaml expects ./tmp/protovalidate/proto/protovalidate
mkdir -p "${PROTO_TMP_DIR}/protovalidate/proto"
ln -sf "${PROTOVALIDATE_DIR}" "${PROTO_TMP_DIR}/protovalidate/proto/protovalidate"

# -------------------------------------------------------------------
# Run Go proto generation
# -------------------------------------------------------------------
echo "→ Running buf dep update..."
buf dep update || echo "⚠ buf dep update had warnings (non-fatal)"
echo "→ Running buf generate (Go)..."
buf generate

# -------------------------------------------------------------------
# Run JS proto generation
# -------------------------------------------------------------------
echo "→ Generating JS..."
cd js
./generate.sh
cd ..

echo "✓ Done"
