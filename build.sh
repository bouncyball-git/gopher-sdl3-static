#!/bin/bash
set -e

MODE=${1:-release}

case "$MODE" in
  debug)
    echo "Building hello (debug)..."
    go build -o hello ./examples/hello/
    echo "built: ./hello (debug)"
    ;;
  release)
    echo "Building hello (release)..."
    go build -ldflags="-s -w" -trimpath -o hello ./examples/hello/
    echo "built: ./hello (release)"
    ;;
  *)
    echo "Usage: $0 [debug|release]"
    exit 1
    ;;
esac
