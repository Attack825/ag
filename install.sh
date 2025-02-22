#!/bin/bash
# 自动检测系统类型
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# 转换架构命名
case $ARCH in
  x86_64) ARCH="amd64" ;;
  aarch64) ARCH="arm64" ;;
esac

# 构建下载URL
DOWNLOAD_URL="https://github.com/irorange27/ag/releases/latest/download/ag-$OS-$ARCH"

# 下载并安装
echo "正在安装 ag for $OS/$ARCH..."
curl -sfL $DOWNLOAD_URL -o /tmp/ag
chmod +x /tmp/ag
sudo mv /tmp/ag /usr/local/bin/ag

# 验证安装
ag --version