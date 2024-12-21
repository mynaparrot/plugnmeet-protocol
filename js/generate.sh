#!/usr/bin/env sh

pnpm install
rm -rf src/gen
buf generate
pnpm run build
