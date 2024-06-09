#!/bin/bash

mkdir -p cache
mkdir -p masterdata

./hailstorm
if [ -f "cache/updated" ]; then
  python3 unpack.py
fi
