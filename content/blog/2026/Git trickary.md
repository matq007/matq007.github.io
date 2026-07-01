---
title: Untitled
date: 2026-04-22T08:00:00
tags:
  - blog/notes
---
This is a list of git tricks I've found randomly on the Internet that could be helpful.

## Amend any commit hash with

```bash
# .gitignore
[alias]
  fixup = "!f() { TARGET=$(git rev-parse $1); git commit --fixup=$TARGET ${@:2} && GIT_SEQUENCE_EDITOR=true git rebase -i --autostash --autosquash $TARGET^; }; f"
```

running `git fixup <SHA-1>`.

## Repo analysis

Most 20 changes files in last year 

```bash
git log --format=format: --name-only --since="1 year ago" | sort | uniq -c | sort -nr | head -20
```

Number of commits per developer

```bash
git shortlog -sn --no-merges
```

Which files are broken the most

```bash
git log -i -E --grep="fix|bug|broken" --name-only --format='' | sort | uniq -c | sort -nr | head -20
```

## References

- [Git fixup is magic (and Magit is too)](https://arialdomartini.github.io/git-fixup)
- [The Git Commands I Run Before Reading Any Code](https://piechowski.io/post/git-commands-before-reading-code/)