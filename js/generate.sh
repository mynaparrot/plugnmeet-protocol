#!/usr/bin/env sh
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

pnpm install
rm -rf src/gen
pnpm exec buf generate
pnpm run build
