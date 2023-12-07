#!/usr/bin/env bash
set -e
on_error() {
  echo "[ERROR] ${BASH_SOURCE[0]}:${BASH_LINENO[0]}: '${BASH_COMMAND}' failed">&2
}
trap on_error ERR

SOURCE_DIR=("./custom" "./ent")
TARGET_DIR="./integrate"

# Create target directory if not exists
mkdir -p "$TARGET_DIR"


for src in "${SOURCE_DIR[@]}"; do
  cp "$src"/*.sql $TARGET_DIR/
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
  fi
done

atlas migrate hash \
  --dir "file://$TARGET_DIR"

echo "integrate_migration.sh done"