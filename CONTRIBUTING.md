# 贡献指南

欢迎来到 AG 项目！我们非常欢迎和感谢您的贡献。以下是参与贡献的指南。

## 如何贡献

### 报告问题

- 在提交 issue 前，请先搜索是否已有类似问题
- 使用清晰、简洁的语言描述问题
- 如果可能，提供复现步骤和预期行为

### 提交 Pull Request

1. Fork 本仓库
2. 创建新的分支：

   ```bash
   git checkout -b feature/your-feature-name
   ```

3. 提交代码变更
4. 推送分支到你的 fork：

   ```bash
   git push origin feature/your-feature-name
   ```

5. 创建 Pull Request

### 代码风格

- 遵循 Go 官方代码风格
- 使用 `gofmt` 格式化代码
- 保持代码简洁、可读

### 提交信息规范

- 使用英文撰写提交信息
- 遵循 Conventional Commits 规范：

  ```plaintext
  <type>[optional scope]: <description>

  [optional body]

  [optional footer(s)]
  ```

  示例：

  ```plaintext
  feat(api): add new provider interface

  - Add Provider interface
  - Implement VolcEngine client
  ```

## 开发环境设置

1. 安装 Go 1.21+
2. 克隆仓库：

   ```bash
   git clone https://github.com/irorange27/ag.git
   cd ag
   ```

3. 安装依赖：

   ```bash
   go mod download
   ```

4. 运行测试：

   ```bash
   go test ./...
   ```

## 代码结构

```plaintext
.
├── api/               # 提供商接口实现
├── cmd/               # 命令行接口
├── config/            # 配置管理
├── internal/          # 内部实现
├── main.go            # 程序入口
├── go.mod             # Go 模块文件
└── go.sum             # 依赖校验文件
```

## 测试

- 添加新功能时请同时添加测试用例
- 运行所有测试：

  ```bash
  go test -v ./...
  ```

- 检查测试覆盖率：

  ```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out
  ```

## 行为准则

请遵守 [贡献者公约](https://www.contributor-covenant.org/version/2/1/code_of_conduct/)。

## 联系方式

如有任何问题，请通过以下方式联系：

- 提交 issue
- 邮件：[orange27\[at\]shu.edu.cn](orange27@shu.edu.cn)
