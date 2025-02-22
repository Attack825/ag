#!/bin/bash
# 编译全平台版本
GOOS_LIST=(darwin linux windows freebsd)
GOARCH_LIST=(amd64 arm64)

for GOOS in "${GOOS_LIST[@]}"; do
  for GOARCH in "${GOARCH_LIST[@]}"; do
    # 压缩符号表减小体积
    OUTPUT="bin/ag-$GOOS-$GOARCH"
    if [ $GOOS = "windows" ]; then
      OUTPUT+=".exe"
    fi
  
    CMD="env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags '-s -w' -o $OUTPUT"
    echo "编译: $OUTPUT"
    $CMD
  done
done

# 进一步压缩（可选）
upx --best bin/ag-*