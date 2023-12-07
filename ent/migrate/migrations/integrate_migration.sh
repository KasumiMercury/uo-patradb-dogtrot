#!/usr/bin/env bash

SOURCE_DIR=("./custom" "./ent")
TARGET_DIR="./integrate"

# Create target directory if not exists
if [ ! -d "$TARGET_DIR" ]; then
  mkdir -p "$TARGET_DIR"
  if [ $? -ne 0 ]; then
    echo "mkdir -p $TARGET_DIR failed"
    exit 1
  fi
fi

for src in "${SOURCE_DIR[@]}"; do
  cp "$src"/*.sql $TARGET_DIR/
  if [ $? -ne 0 ]; then
    echo "cp $src/*.sql $TARGET_DIR/ failed"
    exit 1
  fi
done

for file in "${TARGET_DIR}"/*.sql; do
  filename=$(basename -- "$file")
  found=false

  for src in "${SOURCE_DIR[@]}"; do
    if [ -f "$src/$filename" ]; then
      found=true
      break
    fi
  done

  if [ "$found" = false ]; then
    echo "rm $file"
    rm -f "$file"
    if [ $? -ne 0 ]; then
      echo "rm $file failed"
      exit 1
    fi
  fi
done

atlas migrate hash \
  --dir "file://$TARGET_DIR"
if [ $? -ne 0 ]; then
  echo "atlas migrate hash failed"
  exit 1
fi

echo "integrate_migration.sh done"