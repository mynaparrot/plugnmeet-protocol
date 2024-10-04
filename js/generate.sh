#!/usr/bin/env sh

pnpm install
buf generate
pnpm run build
