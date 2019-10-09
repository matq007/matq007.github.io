---
layout: post
title: "Installing RNA velocity on macOS"
date: 2019-10-09
tags: [bioinformatics, R, rna-velocity]
layout: "post"
---

Package repo: https://github.com/velocyto-team/velocyto.R

At first install necessary library dependencies.

```bash
brew install gcc libmp boost
```

Update your `~/.R/Makevars`. It should look similar to the one below.

```bash
VER=-9
CC=gcc$(VER)
CXX=g++$(VER)
CFLAGS=-mtune=native -g -O2 -Wall -pedantic -Wconversion
CXXFLAGS=-mtune=native -g -O2 -Wall -pedantic -Wconversion
FLIBS=-L/usr/local/Cellar/gcc/9.2.0/lib/gcc/9/

# This is for open mp
SHLIB_OPENMP_CFLAGS=-Xpreprocessor -fopenmp
SHLIB_OPENMP_CXXFLAGS=-Xpreprocessor -fopenmp
```

Additional required R packages.

```R
BiocManager::install(c("pcaMethods", "igraph", "h5"))
install.packages("devtools")
devtools::install_github("velocyto-team/velocyto.R")
```
