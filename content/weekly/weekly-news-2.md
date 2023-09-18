---
title: "Weekly news #2"
date: 2021-01-10
tags: [weekly]
---

## Science/Bioinformatics

### [DeepCell Kiosk: scaling deep learning–enabled cellular image analysis with Kubernetes](https://www.nature.com/articles/s41592-020-01023-0)

_Bannon et al., Nat Methods, 2021_

<a href="#" class="badge badge-primary">#image-analysis</a>
<a href="#" class="badge badge-primary">#kubernetes</a>
<a href="#" class="badge badge-primary">#google-cloud</a>

> 📋New software `DeepCell Kiosk` is designed to scale predefined
> deep-learning image analysis on Google cloud using Kubernetes.

Deep learning is transforming the analysis of biological images, but applying these models to large datasets remains challenging. Here we describe the DeepCell Kiosk, cloud-native software that dynamically scales deep learning workflows to accommodate large imaging datasets. To demonstrate the scalability and affordability of this software, we identified cell nuclei in 106 1-megapixel images in ~5.5 h for ~US$250, with a cost below US$100 achievable depending on cluster configuration. The DeepCell Kiosk can be downloaded at
[https://github.com/vanvalenlab/kiosk-console](https://github.com/vanvalenlab/kiosk-console); a persistent deployment is available at
[https://deepcell.org/](https://deepcell.org/).

### [Spatiotemporal analysis of human intestinal development at single-cell resolution](https://www.cell.com/cell/fulltext/S0092-8674(20)31686-X)

_Fawkner-Corbett et al., Cell, 2021_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#spatial-transcriptomics</a>
<a href="#" class="badge badge-primary">#atlas</a>

> 📋Multimodal atlas of intestinal development.

Development of the human intestine is not well understood. Here, we link single-cell RNA sequencing and spatial transcriptomics to characterize intestinal morphogenesis through time. We identify 101 cell states including epithelial and mesenchymal progenitor populations and programs linked to key morphogenetic milestones. We describe principles of crypt-villus axis formation; neural, vascular, mesenchymal morphogenesis, and immune population of the developing gut. We identify the differentiation hierarchies of developing fibroblast and myofibroblast subtypes and describe diverse functions for these including as vascular niche cells. We pinpoint the origins of Peyer’s patches and gut-associated lymphoid tissue (GALT) and describe location-specific immune programs. We use our resource to present an unbiased analysis of morphogen gradients that direct sequential waves of cellular differentiation and define cells and locations linked to rare developmental intestinal disorders. We compile a publicly available online resource, spatio-temporal analysis resource of fetal intestinal development (STAR-FINDer), to facilitate further work.

### [scMC learns biological variation through the alignment of multiple single-cell genomics datasets](https://genomebiology.biomedcentral.com/articles/10.1186/s13059-020-02238-2)

_Zhang et al., Genome Biol, 2021_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#batch-correction</a>
<a href="#" class="badge badge-primary">#seurat</a>

> 📋Novel scRNA-seq batch correction method which can be easily plugged into your Seurat analysis.

Distinguishing biological from technical variation is crucial when integrating and comparing single-cell genomics datasets across different experiments. Existing methods lack the capability in explicitly distinguishing these two variations, often leading to the removal of both variations. Here, we present an integration method scMC to remove the technical variation while preserving the intrinsic biological variation. scMC learns biological variation via variance analysis to subtract technical variation inferred in an unsupervised manner. Application of scMC to both simulated and real datasets from single-cell RNA-seq and ATAC-seq experiments demonstrates its capability of detecting context-shared and context-specific biological signals via accurate alignment.

### [Control of osteoblast regeneration by a train of Erk activity waves](https://www.nature.com/articles/s41586-020-03085-8)

_De Simone et al., Nature, 2021_

<a href="#" class="badge badge-primary">#live-imaging</a>
<a href="#" class="badge badge-primary">#signaling</a>
<a href="#" class="badge badge-primary">#regeneration</a>
<a href="#" class="badge badge-primary">#matlab</a>

> 📋Insights into how Erk signaling regulates osteoblasts regeneration in zebrafish.

Regeneration is a complex chain of events that restores a tissue to its original size and shape. The tissue-wide coordination of cellular dynamics that is needed for proper morphogenesis is challenged by the large dimensions of regenerating body parts. Feedback mechanisms in biochemical pathways can provide effective communication across great distances, but how they might regulate growth during tissue regeneration is unresolved. Here we report that rhythmic travelling waves of Erk activity control the growth of bone in time and space in regenerating zebrafish scales, millimetre-sized discs of protective body armour. We find that waves of Erk activity travel across the osteoblast population as expanding concentric rings that are broadcast from a central source, inducing ring-like patterns of tissue growth. Using a combination of theoretical and experimental analyses, we show that Erk activity propagates as excitable trigger waves that are able to traverse the entire scale in approximately two days and that the frequency of wave generation controls the rate of scale regeneration. Furthermore, the periodic induction of synchronous, tissue-wide activation of Erk in place of travelling waves impairs tissue growth, which indicates that wave-distributed Erk activation is key to regeneration. Our findings reveal trigger waves as a regulatory strategy to coordinate cell behaviour and instruct tissue form during regeneration.

### [Error correction enables use of Oxford Nanopore technology for reference-free transcriptome analysis](https://www.nature.com/articles/s41467-020-20340-8)

_Sahlin et al., Nat Commun, 2021_

<a href="#" class="badge badge-primary">#long-read</a>
<a href="#" class="badge badge-primary">#nano-pore</a>
<a href="#" class="badge badge-primary">#error-correction</a>

> 📋Nano-pore long-read technology suffers from high sequencing error compared to
> short-read methods. In this paper the authors developed ad new error correction algorithm `IsONcorrect` to tackle this problem.

Oxford Nanopore (ONT) is a leading long-read technology which has been revolutionizing transcriptome analysis through its capacity to sequence the majority of transcripts from end-to-end. This has greatly increased our ability to study the diversity of transcription mechanisms such as transcription initiation, termination, and alternative splicing. However, ONT still suffers from high error rates which have thus far limited its scope to reference-based analyses. When a reference is not available or is not a viable option due to reference-bias, error correction is a crucial step towards the reconstruction of the sequenced transcripts and downstream sequence analysis of transcripts. In this paper, we present a novel computational method to error correct ONT cDNA sequencing data, called isONcorrect. IsONcorrect is able to jointly use all isoforms from a gene during error correction, thereby allowing it to correct reads at low sequencing depths. We are able to obtain a median accuracy of 98.9–99.6%, demonstrating the feasibility of applying cost-effective cDNA full transcript length sequencing for reference-free transcriptome analysis.

## Programming

### [The Little Book of Python Anti-Patterns](https://docs.quantifiedcode.com/python-anti-patterns/index.html)

Do you use python for your work or research? Then you should pay attention
to following some best practices. The one big mistake I see constantly is not
using `@property` in classes. I'm not going to lie, I did this mistake as well,
especially when you are accustomed to develop in different languages.

### [Nerfie](https://nerfies.github.io/)

New approach on how to fully reconstruct a scene from different angles using
your phone. Look at this short [video](https://www.youtube.com/watch?v=IDMiMKWucaI).

## Tools

* [dataframehq/whale](https://github.com/dataframehq/whale): is a lightweight
  data discovery, documentation, and quality engine for your data warehouse
* [JupyterLab 3](https://blog.jupyter.org/jupyterlab-3-0-is-out-4f58385e25bb): time
  for an upgrade 🎉

## Tutorials

* [Advent of Code 2020 in Python by Peter Norvig](https://github.com/norvig/pytudes/blob/master/ipynb/Advent-2020.ipynb)
* [Show progress in your Python apps with tqdm](https://opensource.com/article/20/12/tqdm-python)
* [How to build beautiful plots with Python and Seaborn](https://livecodestream.dev/post/how-to-build-beautiful-plots-with-python-and-seaborn/)

## Others

* [dd, bs= and why you should use conv=fsync](https://abbbi.github.io/dd/)
* [Scipy Lecture Notes (advanced)](http://scipy-lectures.org/index.html)

## Did you know?

### [People With This Rare Brain Condition Can't 'Count Sheep' in Their Mind](https://www.sciencealert.com/people-with-aphantasia-can-t-visualise-sheep-jumping-over-a-fence)

Do you have trouble counting sheep when you are trying to go to sleep? If so,
it means that you might suffer from aphantasia, which was discovered in 19th
century. It's an inability to visualize a mental picture from memory.
