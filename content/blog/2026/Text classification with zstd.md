---
title: "[Weblink]: Text classification with zstd"
date: 2026-02-15T08:00:00
tags:
  - blog/weblink
---
[Max Halford](https://maxhalford.github.io/blog/text-classification-zstd/) showcases how Zstandard compression (zstd) library from Facebook can be used for text classification. This is achieved because the zstd supports incremental compression. The trick Max shows is that one can build multiple topic classifiers and the one returning the smallest size when using test set, is the right one. The module (`compression.zstd`) was introduced in Python 3.14. There is an interesting [discussion on HN](https://news.ycombinator.com/item?id=46942864) about limitations of this approach for AI.