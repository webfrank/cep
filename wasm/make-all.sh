#!/bin/bash

original_dir="$(pwd)"

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

for dir in "$script_dir"/*/; do
  if [ -d "$dir" ]; then
    echo "Making: $dir"
    cd $dir
    make build
  fi
done

cd $original_dir
echo "Done."