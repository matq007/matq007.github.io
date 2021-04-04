---
layout: weekly
title: "Bi-Weekly news #13 and #14"
date: 2021-04-04
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [Richard Stallman is Back!](https://news.itsfoss.com/richard-stallman-is-back-at-fsf/)
* [Gnome 40 is out with some major UI changes](https://forty.gnome.org/)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Abel Prize celebrates union of mathematics and computer science](https://www.nature.com/articles/d41586-021-00694-9)

_Davide Castelvecchi_

<a href="#" class="badge badge-primary">#abel-prize</a>
<a href="#" class="badge badge-primary">#mathematics</a>

_Two pioneers of the theory of computation have won the 2021 Abel Prize, one of the most prestigious honours in mathematics._

### [An Overview of Next-Generation Sequencing](https://www.technologynetworks.com/genomics/articles/an-overview-of-next-generation-sequencing-346532)

_Athina Gkazi, PhD_

<a href="#" class="badge badge-primary">#ngs</a>
<a href="#" class="badge badge-primary">#overview</a>

Over the last 56 years, researchers have been developing methods and technologies to assist in the determination of nucleic acid sequences in biological samples. Our ability to sequence DNA and RNA accurately has had a great impact in numerous research fields. This article discusses what next-generation sequencing (NGS) is, advances in the technology and its applications.

### [Robust integration of multiple single-cell RNA sequencing datasets using a single reference space](https://www.nature.com/articles/s41587-021-00859-x)

_Liu et al., Nat Biotechnol (2021)_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#integration</a>

_In many biological applications of single-cell RNA sequencing (scRNA-seq), an integrated analysis of data from multiple batches or studies is necessary. Current methods typically achieve integration using shared cell types or covariance correlation between datasets, which can distort biological signals. Here we introduce an algorithm that uses the gene eigenvectors from a reference dataset to establish a global frame for integration. Using simulated and real datasets, we demonstrate that this approach, called Reference Principal Component Integration (RPCI), consistently outperforms other methods by multiple metrics, with clear advantages in preserving genuine cross-sample gene expression differences in matching cell types, such as those present in cells at distinct developmental stages or in perturbated versus control studies. Moreover, RPCI maintains this robust performance when multiple datasets are integrated. Finally, we applied RPCI to scRNA-seq data for mouse gut endoderm development and revealed temporal emergence of genetic programs helping establish the anterior–posterior axis in visceral endoderm._

### [Integrated intra‐ and intercellular signaling knowledge for multicellular omics analysis](https://www.embopress.org/doi/full/10.15252/msb.20209923)

_Turei et al., Mol Sys Biol (2021)_

<a href="#" class="badge badge-primary">#signaling</a>
<a href="#" class="badge badge-primary">#cell-interactions</a>

_Molecular knowledge of biological processes is a cornerstone in omics data analysis. Applied to single‐cell data, such analyses provide mechanistic insights into individual cells and their interactions. However, knowledge of intercellular communication is scarce, scattered across resources, and not linked to intracellular processes. To address this gap, we combined over 100 resources covering interactions and roles of proteins in inter‐ and intracellular signaling, as well as transcriptional and post‐transcriptional regulation. We added protein complex information and annotations on function, localization, and role in diseases for each protein. The resource is available for human, and via homology translation for mouse and rat. The data are accessible via OmniPath’s web service (https://omnipathdb.org/), a Cytoscape plug‐in, and packages in R/Bioconductor and Python, providing access options for computational and experimental scientists. We created workflows with tutorials to facilitate the analysis of cell–cell interactions and affected downstream intracellular signaling processes. OmniPath provides a single access point to knowledge spanning intra‐ and intercellular processes for data analysis, as we demonstrate in applications studying SARS‐CoV‐2 infection and ulcerative colitis._

### [sepal: identifying transcript profiles with spatial patterns by diffusion-based modeling](https://academic.oup.com/bioinformatics/advance-article/doi/10.1093/bioinformatics/btab164/6168120)

_Andersson et al., Bioinformatics (2021)_

<a href="#" class="badge badge-primary">#spatial-transcriptomics</a>

_Collection of spatial signals in large numbers has become a routine task in multiple omics-fields, but parsing of these rich datasets still pose certain challenges. In whole or near-full transcriptome spatial techniques, spurious expression profiles are intermixed with those exhibiting an organized structure. To distinguish profiles with spatial patterns from the background noise, a metric that enables quantification of spatial structure is desirable. Current methods designed for similar purposes tend to be built around a framework of statistical hypothesis testing, hence we were compelled to explore a fundamentally different strategy. We propose an unexplored approach to analyze spatial transcriptomics data, simulating diffusion of individual transcripts to extract genes with spatial patterns. The method performed as expected when presented with synthetic data. When applied to real data, it identified genes with distinct spatial profiles, involved in key biological processes or characteristic for certain cell types. Compared to existing methods, ours seemed to be less informed by the genes’ expression levels and showed better time performance when run with multiple cores._

### [User-friendly, scalable tools and workflows for single-cell RNA-seq analysis](https://www.nature.com/articles/s41592-021-01102-w)

_Moreno et al., Nat Methods (2021)_

<a href="#" class="badge badge-primary">#portal</a>
<a href="#" class="badge badge-primary">#scrna-seq</a>

_As single-cell RNA sequencing (scRNA-seq) becomes widespread, accessible and scalable computational pipelines for data analysis are needed. We introduce an interactive computational environment for single-cell studies based on Galaxy1, with functions from established workflows. Single Cell Interactive Application (SCiAp) provides easy access to data from the Human Cell Atlas (HCA) and EMBL-EBI’s Single Cell Expression Atlas (SCEA)2 projects and can be deployed on different computing platforms, making single-cell data analysis of large-scale projects accessible to the scientific community._

### [Enabling high‐throughput biology with flexible open‐source automation](https://www.embopress.org/doi/full/10.15252/msb.20209942)

_Chory et al., Mol Sys Biol (2021)_

<a href="#" class="badge badge-primary">#automation</a>
<a href="#" class="badge badge-primary">#lab</a>

_Our understanding of complex living systems is limited by our capacity to perform experiments in high throughput. While robotic systems have automated many traditional hand‐pipetting protocols, software limitations have precluded more advanced maneuvers required to manipulate, maintain, and monitor hundreds of experiments in parallel. Here, we present Pyhamilton, an open‐source Python platform that can execute complex pipetting patterns required for custom high‐throughput experiments such as the simulation of metapopulation dynamics. With an integrated plate reader, we maintain nearly 500 remotely monitored bacterial cultures in log‐phase growth for days without user intervention by taking regular density measurements to adjust the robotic method in real‐time. Using these capabilities, we systematically optimize bioreactor protein production by monitoring the fluorescent protein expression and growth rates of a hundred different continuous culture conditions in triplicate to comprehensively sample the carbon, nitrogen, and phosphorus fitness landscape. Our results demonstrate that flexible software can empower existing hardware to enable new types and scales of experiments, empowering areas from biomanufacturing to fundamental biology._

### [CytoTree: an R/Bioconductor package for analysis and visualization of flow and mass cytometry data](https://bmcbioinformatics.biomedcentral.com/articles/10.1186/s12859-021-04054-2)

_Dai et al., BMC Bioinformatics (2021)_

<a href="#" class="badge badge-primary">#cytometry</a>

_The rapidly increasing dimensionality and throughput of flow and mass cytometry data necessitate new bioinformatics tools for analysis and interpretation, and the recently emerging single-cell-based algorithms provide a powerful strategy to meet this challenge. Here, we present CytoTree, an R/Bioconductor package designed to analyze and interpret multidimensional flow and mass cytometry data. CytoTree provides multiple computational functionalities that integrate most of the commonly used techniques in unsupervised clustering and dimensionality reduction and, more importantly, support the construction of a tree-shaped trajectory based on the minimum spanning tree algorithm. A graph-based algorithm is also implemented to estimate the pseudotime and infer intermediate-state cells. We apply CytoTree to several examples of mass cytometry and time-course flow cytometry data on heterogeneity-based cytology and differentiation/reprogramming experiments to illustrate the practical utility achieved in a fast and convenient manner. CytoTree represents a versatile tool for analyzing multidimensional flow and mass cytometry data and to producing heuristic results for trajectory construction and pseudotime estimation in an integrated workflow._

---

## <i class="far fa-keyboard"></i> Programming

### [How we built our Python Client that's mostly Rust](https://www.fluvio.io/blog/2021/03/python-client/)

In this post, we’ll talk about how we were able to leverage some of the great Rust tools to build a Python client without writing much Python itself.

### [How to implement a hash table (in C)](https://benhoyt.com/writings/hash-table-in-c/)

Summary: An explanation of how to implement a simple hash table data structure using the C programming language.
I briefly demonstrate linear and binary search, and then design and implement a hash table.
My goal is to show that hash table internals are not scary, but – within certain constraints – are easy enough to build from scratch.

### [C++ in Python the Easy Way!](https://www.youtube.com/watch?v=_5T70cAXDJ0)

Youtube tutorial on how to wrap C++ in python using pybind.

---

## <i class="fas fa-toolbox"></i> Tools

### [Vitals: a tiny macOS process monitor](https://hmarr.com/blog/vitals/)

Vitals is a little Mac app I built recently that shows you which programs are slowing your computer down. It lives in the menu bar, keeping track of resource usage in the background, so you can summon it instantly at the click of a button.

### [wang0618/PyWebIO](https://github.com/wang0618/PyWebIO)

Write interactive web app in script way.

### [synek/git-plan](https://github.com/synek/git-plan)

Git Plan - a better workflow for git

### [Textpattern](https://textpattern.com/)

The small content management system that can handle big ideas

### [Create your own VS Code themes](https://themes.vscode.one/)

Hey! My name is Mike, and I built the VS Code theme studio so you could easily design, deploy, and share your own VS Code themes from scratch. Enjoy!

### [HextEdit](https://hextedit.app/)

A fast and native hex editor for macOS

### [gruns/icecream](https://github.com/gruns/icecream)

🍦 Never use print() to debug again.

### [benfred/py-spy](https://github.com/benfred/py-spy)

Sampling profiler for Python programs

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [My Python testing style guide](https://blog.thea.codes/my-python-testing-style-guide/)

Stargirl Flowers' python practices for testing.

### [SQLite is not a toy database](https://antonz.org/sqlite-is-not-a-toy-database/)

Whether you are a developer, data analyst, QA engineer, DevOps person, or product manager - SQLite is a perfect tool for you. Here is why.

### [socat](https://copyconstruct.medium.com/socat-29453e9fc8a6)

socat stands for SOcket CAT. It is a utility for data transfer between two addresses.

### [The worst so-called “best practice” for Docker](https://pythonspeed.com/articles/security-updates-in-docker/)

Think about switching to non-root in your container.

### [SDS 375](https://wilkelab.org/SDS375/)

Data Visualization in R

---

## <i class="fas fa-rss"></i> Others

### [The Fastest Way to Loop in Python - An Unfortunate Truth](https://www.youtube.com/watch?v=Qgevy75co8c)

What is faster, a for loop, a while loop, or something else? In this video, we try several different ways to accomplish a looping task and discover which is fastest.

### [SSH Keygen: a 2021 update](https://blog.rebased.pl/2021/03/23/ssh-keygen-2021.html)

Read more about new feature of signing your SSH key.

### [Medium Tells Journalists to Feel Free to Quit After Busting Union Drive](https://www.vice.com/en/article/5dp7y3/medium-tells-journalists-to-feel-free-to-quit-after-busting-union-drive)

After what workers describe as a successful union-busting campaign, Ev Williams has announced to journalists who work for him that they should feel free to go.

### [Fundamentals of Optimal Code Style](https://optimal-codestyle.github.io/)

Cognitively-Oriented Approach to Improving Program Readability by _Aleksandr Skobelev_

### [Building a full-text search engine in 150 lines of Python code](https://bart.degoe.de/building-a-full-text-search-engine-150-lines-of-code/)

**Personal recommendation**

---

## <i class="fas fa-rss"></i> Others

### [OpenCell](https://opencell.czbiohub.org/)

Proteome-scale measurements of human protein localization and interactions

### [Study shows survival mechanism for cells under stress](https://www.kth.se/en/aktuellt/nyheter/study-shows-survival-mechanism-for-cells-under-stress-1.1061825)

New research reveals how cancer cells endure stress and survive. Publishing in Molecular Cell, an international research team identified mechanisms that human and mouse cells use to survive heat shock and resume their original function - and even pass the memory of the experience of stress down to their daughter cells.

### [‘This is perhaps the first organism whose evolutionary history was in a computer’](https://thebiologist.rsb.org.uk/biologist-features/professor-michael-levin-interview)

Michael Levin’s work explores how individual cells co-ordinate into multicellular shapes. As Tom Ireland finds, this has not only led to the creation of ‘living robots’ made entirely from frog cells, but it could be used to repair birth defects, regenerate damaged tissue and reprogramme tumours.
