xBERT
|  BERT Typle  | L（网络层数）  |  H（隐藏层维度)  | A（Attention 多头个数）  |  Total Parameters  | 使用GPU内存  |
|  ----  | ----  |  ----  | ----  |  ----  | ----  | 
| Bert base  | 12 |  768  | 12  |  12 * 768 * 12 = 110M  | 7G+  |
| Bert large  | 24 |  1024  | 16  |  340M  | 32G+  | 



# [Blogs]
+ Bert源代码解读, https://github.com/DA-southampton/Read_Bert_Code
+ BERT：深度双向预训练语言模型， https://zhuanlan.zhihu.com/p/355262804
+ NLP Paper, https://github.com/changwookjun/nlp-paper#probe
+ 30分钟带你彻底掌握Bert源码(Pytorch)，超详细！！不看后悔！！https://zhuanlan.zhihu.com/p/148062852
+ BERT language model, https://searchenterpriseai.techtarget.com/definition/BERT-language-model
+ HuggingFace-Transformers系列的下游应用, https://www.jianshu.com/p/cdb13530a8fd
+ 美团BERT的探索和实践, https://tech.meituan.com/2019/11/14/nlp-bert-practice.html
+ ALBERT一作蓝振忠：从谷歌离职回到西湖大学，只为打造一个24小时在线的「心理咨询师」, https://baijiahao.baidu.com/s?id=1705968709780204020&wfr=spider&for=pc
+ “瘦身成功”的ALBERT，能取代BERT吗？https://zhuanlan.zhihu.com/p/115178654
+ 预训练BERT，官方代码发布前他们是这样用TensorFlow解决的, https://zhuanlan.zhihu.com/p/48018623
+ cuda10.2怎么选择tensorflow_gpu? https://www.zhihu.com/question/365667275
+ 

# [Books]
+ Getting Started with Google BERT, https://www.packtpub.com/product/getting-started-with-google-bert/9781838821593


# [Papers]
* NLP Paper, https://github.com/changwookjun/nlp-paper#probe
* [BERT: Pre-training of Deep Bidirectional Transformers for Language Understanding  - NAACL 2019)](https://arxiv.org/abs/1810.04805)  
* [ERNIE 2.0: A Continual Pre-training Framework for Language Understanding - arXiv 2019)](https://arxiv.org/abs/1907.12412)  
* [StructBERT: Incorporating Language Structures into Pre-training for Deep Language Understanding - arXiv 2019)](https://arxiv.org/abs/1908.04577)  
* [RoBERTa: A Robustly Optimized BERT Pretraining Approach  - arXiv 2019)](https://arxiv.org/abs/1907.11692)  
* [ALBERT: A Lite BERT for Self-supervised Learning of Language Representations  - arXiv 2019)](https://arxiv.org/abs/1909.11942)  
* [Multi-Task Deep Neural Networks for Natural Language Understanding  - arXiv 2019)](https://arxiv.org/abs/1901.11504)  
* [What does BERT learn about the structure of language?](https://hal.inria.fr/hal-02131630/document) (ACL2019)
* [Analyzing Multi-Head Self-Attention: Specialized Heads Do the Heavy Lifting, the Rest Can Be Pruned](https://arxiv.org/abs/1905.09418) (ACL2019) [[github](https://github.com/lena-voita/the-story-of-heads)]
* [Open Sesame: Getting Inside BERT's Linguistic Knowledge](https://arxiv.org/abs/1906.01698) (ACL2019 WS)
* [Analyzing the Structure of Attention in a Transformer Language Model](https://arxiv.org/abs/1906.04284) (ACL2019 WS)
* [What Does BERT Look At? An Analysis of BERT's Attention](https://arxiv.org/abs/1906.04341) (ACL2019 WS)
* [Do Attention Heads in BERT Track Syntactic Dependencies?](https://arxiv.org/abs/1911.12246)
* [Blackbox meets blackbox: Representational Similarity and Stability Analysis of Neural Language Models and Brains](https://arxiv.org/abs/1906.01539) (ACL2019 WS)
* [Inducing Syntactic Trees from BERT Representations](https://arxiv.org/abs/1906.11511) (ACL2019 WS)
* [A Multiscale Visualization of Attention in the Transformer Model](https://arxiv.org/abs/1906.05714) (ACL2019 Demo)
* [Visualizing and Measuring the Geometry of BERT](https://arxiv.org/abs/1906.02715)
* [How Contextual are Contextualized Word Representations? Comparing the Geometry of BERT, ELMo, and GPT-2 Embeddings](https://arxiv.org/abs/1909.00512) (EMNLP2019) 
* [Are Sixteen Heads Really Better than One?](https://arxiv.org/abs/1905.10650) (NeurIPS2019)
* [On the Validity of Self-Attention as Explanation in Transformer Models](https://arxiv.org/abs/1908.04211)
* [Visualizing and Understanding the Effectiveness of BERT](https://arxiv.org/abs/1908.05620) (EMNLP2019)
* [Attention Interpretability Across NLP Tasks](https://arxiv.org/abs/1909.11218)
* [Revealing the Dark Secrets of BERT](https://arxiv.org/abs/1908.08593) (EMNLP2019)
* [Investigating BERT's Knowledge of Language: Five Analysis Methods with NPIs](https://arxiv.org/abs/1909.02597) (EMNLP2019)
* [The Bottom-up Evolution of Representations in the Transformer: A Study with Machine Translation and Language Modeling Objectives](https://arxiv.org/abs/1909.01380) (EMNLP2019) 
* [A Primer in BERTology: What we know about how BERT works](https://arxiv.org/abs/2002.12327)
* [Do NLP Models Know Numbers? Probing Numeracy in Embeddings](https://arxiv.org/abs/1909.07940) (EMNLP2019)
* [How Does BERT Answer Questions? A Layer-Wise Analysis of Transformer Representations](https://arxiv.org/abs/1909.04925) (CIKM2019)
* [Whatcha lookin' at? DeepLIFTing BERT's Attention in Question Answering](https://arxiv.org/abs/1910.06431)
* [What does BERT Learn from Multiple-Choice Reading Comprehension Datasets?](https://arxiv.org/abs/1910.12391)
* [Calibration of Pre-trained Transformers](https://arxiv.org/abs/2003.07892)
* [exBERT: A Visual Analysis Tool to Explore Learned Representations in Transformers Models](https://arxiv.org/abs/1910.05276) [[github](https://github.com/bhoov/exbert)]  
* [MobileBERT: a Compact Task-Agnostic BERT for Resource-Limited Devices](https://arxiv.org/pdf/2004.02984.pdf) [[github](https://github.com/google-research/google-research/tree/master/mobilebert)]   
* [Measuring and Reducing Gendered Correlations in Pre-trained Models](https://arxiv.org/pdf/2010.06032.pdf)  
* FastBERT, FastBERT: a Self-distilling BERT with Adaptive Inference Time, https://aclanthology.org/2020.acl-main.537/, https://github.com/autoliuweijie/FastBERT, 
* BERT-of-Theseus: Compressing BERT by Progressive Module Replacing, https://www.connectedpapers.com/main/2e27f119e6fcc5477248eb0f4a6abe8d7cf4f6e7/BERTofTheseus-Compressing-BERT-by-Progressive-Module-Replacing/graph
* BERT-of-Theseus: Compressing BERT by Progressive Module Replacing, https://arxiv.org/abs/2002.02925
* DistilBERT, a distilled version of BERT: smaller, faster, cheaper and lighter, https://arxiv.org/abs/1910.01108
* 


# [Codes]
+ This repository is to collect BERT related resources., https://github.com/Jiakui/awesome-bert
+ BERT, https://github.com/google-research/bert
+ NVIDIA Deep Learning Examples for Tensor Cores, https://github.com/NVIDIA/DeepLearningExamples
+ Classify text with BERT, https://www.tensorflow.org/text/tutorials/classify_text_with_bert#setup
+ BERT with torchtext, https://github.com/pytorch/text/tree/main/examples/BERT
+ 使用huggingface的Transformers预训练自己的bert模型+FineTuning, https://blog.csdn.net/qq_26593695/article/details/115338593
+ Ideas from google's bert for language understanding: Pre-train TextCNN， https://github.com/brightmart/bert_language_understanding
+ BERT-of-Theseus: Compressing BERT by Progressive Module Replacing, https://github.com/JetRunner/BERT-of-Theseus
+ 

## [Codes- Text Classification]
+ Keras BERT, https://github.com/CyberZHG/keras-bert
+ BERT Classifier fine-tuning, https://github.com/Yorko/bert-finetuning-catalyst
+ 

## [Codes- Keras]
+ bert4keras, https://github.com/bojone/bert4keras
+ Keras BERT， https://github.com/CyberZHG/keras-bert
+ 

# [Codes- Cloud]
+ Getting started with the built-in BERT algorithm, https://cloud.google.com/ai-platform/training/docs/algorithms/bert-start

## 【Codes- Chinese】
+ Awesome Pretrained Chinese NLP Models， https://github.com/lonePatient/awesome-pretrained-chinese-nlp-models

## 【Codes- English】

## 【Codes- Kazakh】

## 【Codes- Spanish】






