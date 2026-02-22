---
title: Integrating Jan with opencode
date: 2026-02-22T08:00:00
tags:
  - blog/notes
  - llm
---
This weekend I went back to test again [Jan.ai](https://www.jan.ai/) for running LLM models locally on my MacBook. I do like the simple interface and already build in web-search using [exa](https://exa.ai/). From my crude testing I concluded [MLX optimized models](https://huggingface.co/mlx-community) run faster compared to llama.cpp and I wanted to test it further with opencode. Since I haven't found official documentation on how to integrate `mlx-lm` with opencode I'm providing links and settings on how to achieve it.

Run MLX models with `uvx` one-liner ([kconner.com](https://kconner.com/2025/02/17/running-local-llms-with-mlx.html))

```bash
uvx --python 3.12 --isolated --from mlx-lm mlx_lm.server
```

I wanted to reuse already download MLX models with Jan. By default, they are stored in `~/Library/Application Support/Jan/data/`, which can be changed in the settings.[^1] Alternatively, you can just customize the opencode settings `~/.config/opencode/opencode.json` as follows ([awni](https://gist.github.com/awni/93a973a0cf5fb539b2ce1f37ec4a9989?permalink_comment_id=5931997))

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "mlx": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "MLX (local)",
      "options": {
        "baseURL": "http://127.0.0.1:8080/v1"
      },
      "models": {
        "~/Library/Application Support/Jan/data/mlx/models/Jan-v3-4B-base-instruct-8bit": {
          "name": "Jan-v3-4B-base-instruct-8bit"
        }
      }
    }
  }
}
```

Finally, run opencode (copied from [awni](https://gist.github.com/awni/93a973a0cf5fb539b2ce1f37ec4a9989?permalink_comment_id=5931997))

1. Enter `/connect`
2. Type `MLX` and select it
3. For the API key enter `none`
4. Select the model

[^1]: https://github.com/janhq/jan/issues/2739#issuecomment-2059172637