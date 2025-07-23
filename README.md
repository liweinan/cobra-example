# Cobra CLI 示例项目

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Cobra Version](https://img.shields.io/badge/Cobra-1.9.1-green.svg)](https://github.com/spf13/cobra)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

一个基于 [Cobra](https://github.com/spf13/cobra) 库构建的现代 Go CLI 应用程序示例，展示了如何创建功能丰富的命令行工具。

## ✨ 特性

- 🚀 **现代化 CLI 框架**: 基于 Cobra 库，提供强大的命令行解析能力
- 📝 **自动帮助生成**: 内置帮助系统，支持 `-h` 和 `--help` 标志
- 🔧 **灵活的参数处理**: 支持位置参数、标志参数和参数验证
- 🎯 **子命令支持**: 层次化的命令结构，便于组织复杂功能
- 🌍 **国际化支持**: 中文界面，易于理解和使用
- ⚡ **高性能**: 基于 Go 语言，编译后为单一可执行文件
- 🖥️ **跨平台支持**: 支持 Windows、macOS、Linux，包含 Windows 用户体验优化

## 📋 功能列表

| 命令 | 描述 | 示例 |
|------|------|------|
| `version` | 显示应用程序版本信息 | `./cobra-example version` |
| `greet` | 发送个性化问候消息 | `./cobra-example greet 张三 -f` |
| `echo` | 回显消息，支持格式转换 | `./cobra-example echo "Hello" -u -r 3` |

## 🛠️ 安装

### 前置要求

- Go 1.21 或更高版本
- Git

### 快速开始

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd cobra-example
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **构建项目**
   ```bash
   go build -o cobra-example
   ```

4. **验证安装**
   ```bash
   ./cobra-example version
   ```

## 🚀 使用方法

### 查看帮助信息

```bash
# 查看主帮助
./cobra-example --help

# 查看特定命令帮助
./cobra-example greet --help
./cobra-example echo --help
```

### 版本信息

```bash
./cobra-example version
# 输出: cobra-example v1.0.0
```

### 问候功能

```bash
# 基本问候
./cobra-example greet
# 输出: 你好，世界！

# 个性化问候
./cobra-example greet 张三
# 输出: 你好，张三！

# 正式问候
./cobra-example greet 张三 --formal
# 输出: 您好，张三！很高兴见到您。

# 使用短标志
./cobra-example greet 李四 -f
# 输出: 您好，李四！很高兴见到您。
```

### 回显功能

```bash
# 基本回显
./cobra-example echo "Hello World"
# 输出: Hello World

# 转换为大写
./cobra-example echo "Hello World" --uppercase
# 输出: HELLO WORLD

# 重复消息
./cobra-example echo "Hello" --repeat 3
# 输出:
# Hello
# Hello
# Hello

# 组合使用
./cobra-example echo "Hello" -u -r 2
# 输出:
# HELLO
# HELLO
```

## 📁 项目结构

```
cobra-example/
├── main.go              # 主程序文件，包含所有命令逻辑
├── go.mod               # Go 模块定义文件
├── go.sum               # 依赖校验文件
├── README.md            # 项目说明文档
├── LICENSE              # 许可证文件
└── cobra-example        # 编译后的可执行文件
```

### 📦 依赖说明

| 依赖包 | 版本 | 作用 |
|--------|------|------|
| `github.com/spf13/cobra` | v1.9.1 | 核心 CLI 框架 |
| `github.com/spf13/pflag` | v1.0.6 | POSIX 兼容的标志处理 |
| `github.com/inconshreveable/mousetrap` | v1.1.0 | Windows 平台用户体验优化 |

## 🔧 开发指南

### 添加新命令

1. **定义命令结构**
   ```go
   var newCmd = &cobra.Command{
       Use:   "new [arg]",
       Short: "新命令的简短描述",
       Long:  `新命令的详细描述`,
       Args:  cobra.ExactArgs(1),
       Run: func(cmd *cobra.Command, args []string) {
           // 命令逻辑
       },
   }
   ```

2. **添加标志**
   ```go
   newCmd.Flags().StringP("flag", "f", "", "标志描述")
   newCmd.Flags().BoolP("bool-flag", "b", false, "布尔标志描述")
   ```

3. **注册命令**
   ```go
   func init() {
       rootCmd.AddCommand(newCmd)
   }
   ```

### 构建和测试

```bash
# 构建项目
go build -o cobra-example

# 运行测试
go test ./...

# 代码格式化
go fmt ./...

# 代码检查
golangci-lint run
```

## 🎓 学习要点

这个项目展示了 Cobra 的核心概念：

1. **命令层次结构**: 根命令和子命令的组织方式
2. **参数处理**: 位置参数和标志参数的使用
3. **参数验证**: 使用 `cobra.Args` 验证参数
4. **标志系统**: 短标志和长标志的定义
5. **帮助系统**: 自动生成帮助信息
6. **错误处理**: 优雅的错误处理和退出

## 🔧 技术细节

### Mousetrap 机制

本项目使用了 `github.com/inconshreveable/mousetrap` 库，这是一个专门为 Windows 平台设计的用户体验优化工具。

#### 🐭 作用说明

- **检测启动方式**: 检测 CLI 程序是否通过 Windows 资源管理器双击启动
- **防止误操作**: 当用户双击程序时显示帮助信息，引导正确使用
- **跨平台兼容**: 在 macOS/Linux 上不执行任何操作

#### 💻 Windows 上的行为

当 Windows 用户双击 CLI 程序时：

```
This is a command line tool.

You need to open cmd.exe and run it from there.
```

程序会等待 5 秒后自动退出，或等待用户按回车键。

#### ⚙️ 配置选项

```go
// 自定义帮助文本
var MousetrapHelpText = `自定义的帮助信息`

// 控制显示时长（0 表示等待用户按键）
var MousetrapDisplayDuration = 5 * time.Second

// 完全禁用 mousetrap
var MousetrapHelpText = ""
```

#### 🌍 平台兼容性

| 平台 | 行为 |
|------|------|
| Windows | 检测双击启动，显示帮助信息 |
| macOS/Linux | 无特殊行为，正常执行 |

这个机制确保了跨平台 CLI 工具在 Windows 上的良好用户体验！

## 🤝 贡献指南

我们欢迎所有形式的贡献！

### 贡献方式

1. **报告问题**: 在 GitHub Issues 中报告 bug 或提出建议
2. **提交代码**: Fork 项目并提交 Pull Request
3. **改进文档**: 帮助改进 README 和代码注释
4. **分享经验**: 在 Discussions 中分享使用经验

### 开发流程

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Cobra](https://github.com/spf13/cobra) - 强大的 Go CLI 框架
- [Go 语言](https://golang.org/) - 优秀的编程语言
- 所有贡献者和用户

## 📞 联系方式

- 项目主页: [GitHub Repository](https://github.com/your-username/cobra-example)
- 问题反馈: [GitHub Issues](https://github.com/your-username/cobra-example/issues)
- 讨论交流: [GitHub Discussions](https://github.com/your-username/cobra-example/discussions)

---

⭐ 如果这个项目对您有帮助，请给我们一个星标！
