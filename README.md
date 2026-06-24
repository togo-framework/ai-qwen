<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/ai-qwen</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/ai-qwen"><img src="https://pkg.go.dev/badge/github.com/togo-framework/ai-qwen.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/ai-qwen
```

<!-- /togo-header -->

# ai-qwen

Qwen (DashScope) driver for the togo `ai` plugin — a unified LLM interface (chat + embeddings) for togo apps.

## Install
```bash
togo install togo-framework/ai-qwen
```

## Configure
Set `AI_DRIVER=qwen` and:
```env
DASHSCOPE_API_KEY=...
# optional: DASHSCOPE_BASE_URL=https://dashscope.aliyuncs.com/compatible-mode/v1
```

Then the `ai` plugin routes `Chat`/`Embed` through Qwen (DashScope). Token usage is reported via `ai.Usage` (for the billing plugin).

MIT © ToGO

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
