#!/bin/bash
# 编译全平台版本
GOOS_LIST=(darwin linux windows freebsd)
GOARCH_LIST=(amd64 arm64)

# 确保 bin 目录存在
mkdir -p bin

for GOOS in "${GOOS_LIST[@]}"; do
  for GOARCH in "${GOARCH_LIST[@]}"; do
    # 压缩符号表减小体积
    OUTPUT="bin/ag-$GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
      OUTPUT+=".exe"
    fi

    # 使用 eval 来正确解析命令
    CMD="env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags '-s -w' -o $OUTPUT"
    echo "编译: $OUTPUT"
    eval $CMD
  done
done

# 进一步压缩（可选）
if command -v upx &> /dev/null; then
  upx --best bin/ag-*
else
  echo "upx 未安装，跳过进一步压缩。"
fi