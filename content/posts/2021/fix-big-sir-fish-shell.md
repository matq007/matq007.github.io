---
title: Slow autocomplete in fish-shell on macOS Big Sur
summary: When the autocomplete takes for ever.
date: 2021-01-04
categories: ["devops"]
tags: [fish-shell, brew, macos]
---

**UPDATE**: Upgrading to fish >3 solves the issues now.

If you have decided to upgrade to the latest macOS Big Sur `10.14.6`,
you might have experienced an issue with slow autocomplete in `fish-shell`.

Due to change to read-only permissions for `whatis` [^1] the man pages
can't be updated for `/usr/share/man`. Therefore, the fish is switching to
using `apropos` [^2].

## How to fix without updating?

```bash
# create backup
cp /usr/local/share/fish/functions/__fish_describe_command.fish \
   /usr/local/share/fish/functions/__fish_describe_command.fish.old

# Taken from: https://github.com/fish-shell/fish-shell/issues/6270#issuecomment-548515306
echo "function __fish_describe_command; end" > /usr/local/share/fish/functions/__fish_describe_command.fish
```

Thank you `fish` community for such a quick solution [^3].

[^1]: https://github.com/fish-shell/fish-shell/issues/6270
[^2]: https://apple.stackexchange.com/questions/374025/errors-from-whatis-command-unable-to-rebuild-database-with-makewhatis
[^3]: https://github.com/fish-shell/fish-shell/pull/7365