# Pilot — TODO 与路线图

> 终端命令助手。使用自然语言提问，生成或解释命令。

---

## 状态说明

| 符号 | 含义 |
|------|------|
| ✅ | 已完成 |
| 🔄 | 进行中 |
| 🔲 | 计划中 |
| 💡 | 想法 / 研究 |
| ❌ | 已取消 |

---

## ✅ v0.1.0 — MVP（已完成）

- [x] 使用 Go 构建 CLI 架构
- [x] 集成 OpenRouter 免费层（`openrouter/free`）
- [x] `pilot ask` — 自然语言 → 命令生成
- [x] `pilot explain` — 解释命令
- [x] `pilot run` — 生成 + 确认 + 执行
- [x] `pilot setup` — API Key 配置指南
- [x] 内置 cheatsheet（terminal, git, docker）
- [x] 跨平台支持（macOS, Linux, Windows）
- [x] 彩色终端输出（`fatih/color`）
- [x] 初始提交与仓库创建

---

## 🔄 v0.2.0 — 核心体验（进行中）

### 安全
- [ ] 危险命令检测（`rm -rf`, `reset --hard`, `DROP TABLE` 等）
- [ ] 红色 ⚠️ 警告面板
- [ ] `--force` 跳过警告

### CLI 体验
- [ ] 单一可执行文件（`pilot`）
- [ ] 全局安装脚本
- [ ] 交互式 REPL 模式
- [ ] 加载动画（等待 AI 响应）
- [ ] `--dry-run`（仅显示不执行）

### Cheatsheet
- [ ] 添加 npm / yarn
- [ ] 添加 Kubernetes / kubectl
- [ ] 扩展 SSH / SCP

---

## 🔲 v0.3.0 — 历史与记忆

- [ ] 历史记录
- [ ] 搜索历史
- [ ] 清除历史
- [ ] SQLite 存储
- [ ] 常用命令推荐
- [ ] 收藏功能

---

## 🔲 v0.4.0 — 配置与自定义

- [ ] `config.toml`
- [ ] 模型选择
- [ ] 语言选择
- [ ] 自定义 cheatsheet
- [ ] 配置命令
- [ ] 安全存储 API Key

---

## 🔲 v0.5.0 — 本地化

- [ ] 多语言 UI
- [ ] 自动检测系统语言
- [ ] 手动切换

---

## 🔲 v1.0.0 — 生产就绪

### 分发
- [ ] Homebrew
- [ ] apt
- [ ] winget
- [ ] GitHub Releases
- [ ] 安装脚本

### CI/CD
- [ ] GitHub Actions
- [ ] 跨平台构建
- [ ] 发布自动化
- [ ] 覆盖率报告

### 文档
- [ ] README
- [ ] CONTRIBUTING
- [ ] CHANGELOG
- [ ] 演示

---

## 💡 v1.x — 未来方向

- [ ] 多步骤命令链
- [ ] 插件系统
- [ ] 详细解释模式
- [ ] Shell 自动补全
- [ ] 自动更新
- [ ] 离线模式（Ollama）
- [ ] VS Code 插件
- [ ] Web UI
- [ ] 遥测（可选）

---

## 🐛 已知问题

- [ ] 解析 bug
- [ ] Windows 颜色问题
- [ ] 长响应 UI 破坏

---

## 📦 依赖

| 包 | 版本 | 用途 |
|----|------|------|
| `github.com/fatih/color` | latest | 终端颜色输出 |