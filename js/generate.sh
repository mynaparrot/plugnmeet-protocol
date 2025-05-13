#!/usr/bin/env sh

pnpm install
rm -rf src/gen
pnpm exec buf generate
pnpm run build
