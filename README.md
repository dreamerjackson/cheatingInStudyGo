# 第一章:为什么设计这门课程

&emsp;&emsp;2022年开年，带给go语言初学者一个好消息。今年笔者将会联合我的圈中密友，扫地僧级别的辛蔚大佬，倾尽所能，为初学者打造一套免费的Go语言入门课程。\
&emsp;&emsp;去年底笔者靠一本《Go语言底层原理剖析》，成功拿到了电子工业出版社年度优秀作者，CSDN 2021年度图书。沉寂许久之后，我将再度出发，在这汹涌的时代浪潮中奋进...\
&emsp;&emsp;《计算机程序的构造和解释》曾说过这样一段话：在头脑中孵化的想法，都可以通过计算机程序进行表达与建模。我们不断加深理解、完善模型，让程序逼近真实的心理过程。\
&emsp;&emsp;是的，构造计算机程序就像是创造一个有生命的神灵。我们用想象力与创造力描述心中所想，而这一切复杂到难以置信甚至充满智能的程序都是由简单的符号组合而成，这是多么巧妙、多么让人兴奋。\
&emsp;&emsp;每隔一段时间就会有一门新的编程语言诞生，顺应时代的变化，硬件的发展以及计算机科学的进步。这一次我将用特别的方式来教授这门新兴并富有潜力的Go语言。论述如何正确地使用Go语言来完成我们心中所想....\
&emsp;&emsp;希望当完成这门课程，我们学到的不止是工具，还有"智慧"。是搭建起Go知识体系的智慧，是构建大规模复杂程序的智慧，是掌握了进一步学习能力的智慧！

### 为什么设计这门课程

* 目前市面上的入门资料水平参差不一，讲解者的水平和用心程度差别很大。
* 资料碎片化，没有体系化的整理和洞察力的见解。更多的是知识的堆砌，而没有知识之前的联系。
* 很多资料已经过时老旧，描述错误导致的误导。
* 更多是对单一知识点的用法介绍，却少有涉及到如何正确的使用，为什么要这样设计，和其他语言的对比。
* 学习形式单一，缺少启发式的教学方式和衍生进一步学习的参考资料。

为了弥补当前Go入门视频的不足，笔者即将推出这一套基础课程，学习目标是快速全面系统掌握Go用法与生态，会使用并能够正确的使用Go语言。

### 课程内容

如下课程的思维导图，内容主要分为了几个大的部分。 第一部分学习准备，将介绍这门课程的创作背景、如何使用该课程更好地学习Go语言以及Go开发环境的安装和历史。\
第二部分基础语法和语言特性，将介绍构建Go程序时必不可少的结构与要素以及与其他语言的区别。这些要素就像是钉子、锤子等工具，是原始但基础的。\
第三部分将介绍Go语言赖以成名的高并发编程。介绍如何使用轻量级的协程与通道，屏蔽多线程开发带来的心智负担，使用合理的并发模式，快速构建高并发应用。\
第四部分将介绍Go语言构建大规模复杂系统的设计哲学，介绍如何布局代码，管理依赖并使用面向组合的方式构建优雅可维护的项目。\
第五部分将介绍Go生态的细节，除了能够写代码，还需要会调试，会测试，会观测程序的运行指标。另外成为一个成熟开发者的标志是熟练的掌握了Go标准库与优秀的第三方库开源库，做到游刃有余。\
第六部分最后，将介绍Go语言即将和未来将带来的一些新变化，包括万众瞩目的泛型设计，以及如何进一步学习完成进阶。\


![](<.gitbook/assets/image (1) (1).png>)

完成这一系列需要时间、经验、创造力和勇气。更重要的是，需要你的共同参与。

### 课程形式

* 视频 (https://www.bilibili.com/video/BV1Vi4y1f78S?spm\_id\_from=333.999.0.0)
* 文章 (https://github.com/dreamerjackson/cheatingInStudyGo/edit/master/README.md)
* 衍生阅读材料
* 作业 重要资料将完全开源，持续更新，共同参与，并提供良好的学习群氛围。希望开源共同参与，提供全面、详细、清晰、顶级的内容质量和学习方法。

### 目标人群

* Go 语言初学者
* 使用 Go语言 一段时间但是知识未形成体系

### 为什么是这个名字

学习是一个曲折反复的过程，然而我们不得不承认的是，确实有更快更好的学习方式。就好像我们学习了多年的应试教育英语仍然是哑巴一样，有时候取得成功并不都是努力与否的问题，而是需要找到合适的方向和方法。\
我们总是羡慕游戏中的`氪金`玩家，甚至有些人是上天遁地，无所不能的开挂选手。这是我幽默的比喻，如果我们把学习Go语言当做是一场游戏，入门就是游戏的开局。有些人开局只有一条狗，我希望这一套课程就像是你氪金之后得到的葵花宝典，开局就是豪华装备加持，一路上还有高级陪练在你的身边，刷怪升级变成了一种享受，早早就修炼下山了....

——————

文章列表：

* [2.如何学习Go语言](article/2.howToLearn.md)
* [3.搭建Go语言开发环境](article/3.environment.md)
* [4.书写第一个go语言程序](article/4.%20hello-world.md)
* [5.Go命令行工具](article/5.command-tool.md)
* [6.集成开发环境安装](article/6.ide.md)
* [7.变量与类型](article/7.variables.md)
* [8.表达式与运算符](article/8.expression.md)
* [9.程序控制结构](article/9.control-structures.md)

如果你也希望一起参与到这门课程的学习中来，可加笔者微信,进学习群。\


![笔者微信](https://user-images.githubusercontent.com/42735226/153756008-1c9335b5-c0b0-4456-889f-58a42eaca23f.png)
