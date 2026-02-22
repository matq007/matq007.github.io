---
title: Epsilon in AdamW matters
date: 2026-01-24T12:00:00
tags:
  - blog/weblink
  - data-science
  - llm
---
When training LLM, [Sifal](https://sifal.social) shows that switching from default `eps=1e-8` to `eps=1e-10` in AdamW optimizer can lead to better results. He showcases this on his toy example where default epsilon oscillates when searching for local minima compared to proposed one. However, this only applies when training is **NOT** done with [half-precision](https://docs.pytorch.org/docs/stable/notes/numerical_accuracy.html#reduced-precision-reduction-for-fp16-and-bf16-gemms).

> Klioui, S. (2026). The Epsilon Trap: When Adam Stops Being Adam. Sifal Klioui Blog. [link](https://sifal.social/posts/The-Epsilon-Trap-When-Adam-Stops-Being-Adam/)