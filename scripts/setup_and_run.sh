#!/bin/sh

if ! ./scripts/setup_furya.sh; then
  exit 1
fi
./scripts/run_furya.sh

