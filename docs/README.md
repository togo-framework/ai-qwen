# ai-qwen — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package qwen is a Qwen (DashScope) driver for togo ai (OpenAI-compatible API).
Blank-import it and set AI_DRIVER=qwen + DASHSCOPE_API_KEY.

## Install

```bash
togo install togo-framework/ai-qwen
```

Set `AI_DRIVER=qwen`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `DASHSCOPE_BASE_URL` | _see provider docs_ |
| `G` | _see provider docs_ |

## Usage

```go
provider := ai.FromKernel(k)
resp, err := provider.Chat(ctx, []ai.Message{{Role: "user", Content: "Hello"}}, ai.Options{})
// streaming + provider.Embed(ctx, texts) for vectors; resp.Usage carries token counts
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-qwen
- README: ../README.md
