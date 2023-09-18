---
title: "Weekly news #3"
date: 2021-01-17
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* **Nature Computational Science** [issue #1](https://www.nature.com/natcomputsci/volumes/1/issues/1) is out 🎉
* [Pirated Academic Database Sci-Hub Is Now on the ‘Uncensorable Web’](https://www.nasdaq.com/articles/pirated-academic-database-sci-hub-is-now-on-the-uncensorable-web-2021-01-11)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Lessons from the COVID-19 pandemic for advancing computational drug repurposing strategies](https://www.nature.com/articles/s43588-020-00007-6)

_Galindez et al., Nat Comput Sci, 2021_

<a href="#" class="badge badge-primary">#covid-19</a>
<a href="#" class="badge badge-primary">#drug-disvocery</a>

> _Responding quickly to unknown pathogens is crucial to stop uncontrolled spread of diseases that lead to epidemics, such as the novel coronavirus, and to keep protective measures at a level that causes as little social and economic harm as possible. This can be achieved through computational approaches that significantly speed up drug discovery. A powerful approach is to restrict the search to existing drugs through drug repurposing, which can vastly accelerate the usually long approval process. In this Review, we examine a representative set of currently used computational approaches to identify repurposable drugs for COVID-19, as well as their underlying data resources. Furthermore, we compare drug candidates predicted by computational methods to drugs being assessed by clinical trials. Finally, we discuss lessons learned from the reviewed research efforts, including how to successfully connect computational approaches with experimental studies, and propose a unified drug repurposing strategy for better preparedness in the case of future outbreaks._

### [isoCirc catalogs full-length circular RNA isoforms in human transcriptomes](https://www.nature.com/articles/s41467-020-20459-8)

_Xin et al., Nat Commun, 2021_

<a href="#" class="badge badge-primary">#circular-rna</a>
<a href="#" class="badge badge-primary">#sequencing</a>

> _Circular RNAs (circRNAs) have emerged as an important class of functional RNA molecules. Short-read RNA sequencing (RNA-seq) is a widely used strategy to identify circRNAs. However, an inherent limitation of short-read RNA-seq is that it does not experimentally determine the full-length sequences and exact exonic compositions of circRNAs. Here, we report isoCirc, a strategy for sequencing full-length circRNA isoforms, using rolling circle amplification followed by nanopore long-read sequencing. We describe an integrated computational pipeline to reliably characterize full-length circRNA isoforms using isoCirc data. Using isoCirc, we generate a comprehensive catalog of 107,147 full-length circRNA isoforms across 12 human tissues and one human cell line (HEK293), including 40,628 isoforms ≥500 nt in length. We identify widespread alternative splicing events within the internal part of circRNAs, including 720 retained intron events corresponding to a class of exon-intron circRNAs (EIciRNAs). Collectively, isoCirc and the companion dataset provide a useful strategy and resource for studying circRNAs in human transcriptomes._

### [A robust unsupervised machine-learning method to quantify the morphological heterogeneity of cells and nuclei](https://www.nature.com/articles/s41596-020-00432-x)

_Phillip et al., Nat Protoc, 2021_

<a href="#" class="badge badge-primary">#machine-learning</a>
<a href="#" class="badge badge-primary">#morphology</a>

> _Cell morphology encodes essential information on many underlying biological processes. It is commonly used by clinicians and researchers in the study, diagnosis, prognosis, and treatment of human diseases. Quantification of cell morphology has seen tremendous advances in recent years. However, effectively defining morphological shapes and evaluating the extent of morphological heterogeneity within cell populations remain challenging. Here we present a protocol and software for the analysis of cell and nuclear morphology from fluorescence or bright-field images using the VAMPIRE algorithm (https://github.com/kukionfr/VAMPIRE_open). This algorithm enables the profiling and classification of cells into shape modes based on equidistant points along cell and nuclear contours. Examining the distributions of cell morphologies across automatically identified shape modes provides an effective visualization scheme that relates cell shapes to cellular subtypes based on endogenous and exogenous cellular conditions. In addition, these shape mode distributions offer a direct and quantitative way to measure the extent of morphological heterogeneity within cell populations. This protocol is highly automated and fast, with the ability to quantify the morphologies from 2D projections of cells seeded both on 2D substrates or embedded within 3D microenvironments, such as hydrogels and tissues. The complete analysis pipeline can be completed within 60 minutes for a dataset of ~20,000 cells/2,400 images._

---

## <i class="far fa-keyboard"></i> Programming

### [Docker BuildKit: faster builds, new features, and now it’s stable](https://pythonspeed.com/articles/docker-buildkit/)

Speed up your docker image builds by using stable version of `BuildKit` (`Docker >=19.03`).

To enable it run

1. `export DOCKER_BUILDKIT=1`
2. add `# syntax=docker/dockerfile:1.2` at the very beginning of your `Dockerfile`

### [GitHub security features: highlights from 2020](https://github.blog/2021-01-11-github-security-features-highlights-from-2020/)

With New Year comes new features from GitHub like scanning code for security
issues, accidental push of your api keys/tokens or dependency review. Read the
full list at the GitHub blog.

---

## <i class="fas fa-toolbox"></i> Tools

### [workflowr: organized + reproducible + shareable data science in R](https://jdblischak.github.io/workflowr/)

If you are an R user and you love notebooks, you have to check this package out.
It combines knitr/RMarkdown with version control and generates a easily
customizable website with all the results. Next when you share your data, give
it a try. Here is a [demo project](https://wolfemd.github.io/NRCRI_2020GS/index.html).

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [Tracing the Python GIL](https://www.maartenbreddels.com/perf/jupyter/python/tracing/gil/2021/01/14/Tracing-the-Python-GIL.html)

If you have ever written a multi-threading script in python you probably heard about
Global Lock Interpreter (GIL). If not, imagine a scenario where multiple processes
are writing simultaneously data into one file. Who gets to write first? How do we preserve
the proper state of the file and don't overwrite the data? _Maarten Breddels_ wrote
an amazing blog post diving into details of GIL and even wrote a module
which you can easily plug-in to your Jupyter Notebook and visualize the GIL in action.

### [Ten Quick Tips for Deep Learning in Biology](https://github.com/Benjamin-Lee/deep-rules)

This is a **must read** before you dive into the rabbit hole of AI/Deep learning.

### [This is how I did it: created the best reference manager set up for research & writing](https://www.nrel.colostate.edu/set-up-best-reference-manager/)

Perfect tutorial how to take your `Zotero` setup to the next level. I have very
similar setup and it works like a charm!

### [Go crazy with GitHub Actions](https://sanderknape.com/2021/01/go-crazy-github-actions/)

Master the art of `CI/CD` for your application.

---

## <i class="fas fa-rss"></i> Others

### [OpenLearn](https://www.open.edu/openlearn/free-ebooks)

Free online resources for learning.

### [We Downloaded 10,000,000 Jupyter Notebooks From Github – This Is What We Learned](https://blog.jetbrains.com/datalore/2020/12/17/we-downloaded-10-000-000-jupyter-notebooks-from-github-this-is-what-we-learned/)

JetBrains had a look at publicly available notebook and did some reporting on it.

### [An unlikely database migration](https://tailscale.com/blog/an-unlikely-database-migration/)

How TailScale used `etcd` from Kubernetes and adjusted it to their needs.

---

## <i class="far fa-surprise"></i> Did you know?

There are some peculiar spots across the world which seem to defy laws of gravity.
You throw a ball down the hill and it comes back?

![](https://www.sciencealert.com/images/2017-03/ball-hill.gif)

**Spoiler alerts**, it's not a gravity flaw, but apparently a mind trick by our eyes.
Similar phenomena can be also found in Scotland. Measurements of the road showed
that spot we see as top of the hill had a lower height above the sea level than
then bottom one. The illusion was due to the surrounding area. Head over
to [the article](https://www.sciencealert.com/these-gravity-defying-hills-are-one-of-the-weirdest-natural-phenomena-we-ve-seen)
written by Sciencealert to see more.
