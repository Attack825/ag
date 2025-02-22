# AG - AI 命令行工具

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/irorange27/ag)](https://goreportcard.com/report/github.com/irorange27/ag)
[![Build Status](https://github.com/irorange27/ag/actions/workflows/build.yml/badge.svg)](https://github.com/irorange27/ag/actions)

AG 是一个与 AI 模型交互的命令行工具，支持多种 AI 提供商和流式响应。

## 功能特性

- 🚀 支持多种 AI 提供商（VolcEngine, OpenAI 等）
- 💬 交互式聊天模式
- ⚡ 流式响应，实时显示结果
- 🔧 可配置的提供商设置
- 📦 跨平台支持（Windows, Linux, macOS）

## 安装

### 二进制文件

从 [Releases](https://github.com/irorange27/ag/releases) 页面下载预编译的二进制文件。

### 从源码编译

1. 确保已安装 Go 1.20+
2. 克隆仓库：

   ```bash
   git clone https://github.com/irorange27/ag.git
   cd ag
   ```

3. 编译：

   ```bash
   ./build.sh
   ```

   编译后的文件会在 `bin/` 目录下

## 使用说明

1. 创建配置文件 `config.yaml`：

   ```yaml
   default_provider: "volcengine"
   
   providers:
     volcengine:
       type: volcengine
       endpoint: https://ark.cn-beijing.volces.com/api/v3/chat/completions
       api_key: your-api-key-here
       model: deepseek-v3-241226
   ```

2. 启动交互模式：

   ```bash
   ./ag
   ```

3. 单次对话：

   ```bash
   ./ag chat "你好，世界！"
   ```

## 配置

配置文件支持以下选项：

- `default_provider`: 默认使用的 AI 提供商
- `providers`: 提供商配置
  - `type`: 提供商类型（volcengine, openai）
  - `endpoint`: API 地址
  - `api_key`: API 密钥
  - `model`: 使用的模型

## 贡献

欢迎贡献！请阅读 [CONTRIBUTING.md](CONTRIBUTING.md) 了解如何参与开发。

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](LICENSE) 文件。
