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
