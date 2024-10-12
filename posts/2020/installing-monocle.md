---
title: "Installing monocle3"
date: 2019-09-22
tags: [bioinformatics, R, monocle3]
---

In this short post I will demonstrate how to install package `monocle3` in R on
your macOS. If you are using default compilers provided by Apple, you might have
encounter some errors like

```R
clang: error: linker command failed with exit code 1 (use -v to see invocation)
make: *** [leidenbase.so] Error 1 ERROR: compilation failed for package ‘leidenbase’
```

or

```R
ERROR: compilation failed for package ‘leidenbase’ ...
```

## Solution

To solve this issue you first have to update your current `gcc` compiler as well
as installing `llvm`. The problem is that macOS does not come with `gFortran`
package which is required by sub-package `leidenbase`. In this case I recommend
installing/upgrading the packages using [brew](http://brew.sh) illustrated below.

```bash
brew install gcc llvm
# copy and paste the code in your shell to enable llvm
```

Secondly, now we have to tell `R` to use these packages by updating `~/.R/Makevars`
path. If you do not have `.R` folder in your home directory, you can create it 
by running `mkdir ~\.R`.

Before changing the variables you have to verify which `gcc` version have you
installed. You can do so by running `gcc-<TAB>`. The example below shows how
to setup the variables for `gcc-9`.

```bash
vim ~/.R/Makevars
# Copy and paste the code below.
VER=-9
CC=gcc$(VER)
CXX=g++$(VER)
CFLAGS=-mtune=native -g -O2 -Wall -pedantic -Wconversion
CXXFLAGS=-mtune=native -g -O2 -Wall -pedantic -Wconversion
FLIBS=-L/usr/local/Cellar/gcc/9.2.0/lib/gcc/9/

# Restart R
```

Good job, now restart `R` and run the `monocle3` installation again.

```R
BiocManager::install(c('BiocGenerics', 'DelayedArray', 'DelayedMatrixStats',
                       'limma', 'S4Vectors', 'SingleCellExperiment',
                       'SummarizedExperiment'))
devtools::install_github('cole-trapnell-lab/leidenbase')
devtools::install_github('cole-trapnell-lab/monocle3')
```

## Resources

- [Stackoverflow](https://stackoverflow.com/questions/43597632/understanding-the-contents-of-the-makevars-file-in-r-macros-variables-r-ma)
- [Monocle3 documentation](https://cole-trapnell-lab.github.io/monocle3/docs/installation/)
