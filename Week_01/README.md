# 学习笔记

## 前言：

![image-20200830224004795](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20200830224004795.png)

## 训练环境设置、编码技巧和Code Style

- [Windows Microsoft New Terminal](http://github.com/microsoft/terminal)
- [VS Code Themes](http://vscodethemes.com/)
- [教你打造一款颜值逆天的 VS Code](http://toutiao.io/posts/7w5ixl/preview)
- [炫酷的 VS Code 毛玻璃效果](http://juejin.im/post/5ce1365151882525ff28ed47)
- [自顶向下的编程方式](http://markhneedham.com/blog/2008/09/15/clean-code-book-review/)
- [自顶向下编程的 LeetCode 例题](http://leetcode-cn.com/problems/valid-palindrome/)

## 时间和空间复杂度

- [如何理解算法时间复杂度的表示法](http://www.zhihu.com/question/21387264)
- [wiki：Master theorem](http://en.wikipedia.org/wiki/Master_theorem_(analysis_of_algorithms))
- [wiki：主定理](http://zh.wikipedia.org/wiki/主定理)

## 1. 数组，链表，调表

数组：在内存中开辟连续的内存地址，存储元素

链表：当前Node对象存储当前节点值与下一节点内存地址

调表：带有调表索引的链表，只能用于元素有序的情况下，用来取代平衡树二分查找

|          | 左append | 右append | 查询     | 插入     | 删除     |
| -------- | -------- | -------- | -------- | -------- | -------- |
| 数组     | O(1)     | O(1)     | O(1)     | O(n)     | O(n)     |
| 普通链表 | O(1)     | O(1)     | O(n)     | O(1)     | O(1)     |
| 调表     | O(1)     | O(1)     | O(log n) | O(log n) | O(log n) |

空间复杂度上，数组最少，普通链表第二，跳表最高，但，都是O(n)。

链表应用：LRU Cache

跳表应用：Redis

## 2.栈和队列

栈（stack）：先入后出（LIFO/FILO）

队列(Queue)：先入先出（FIFO）

|		| 左append | 右append | 查询 | 插入 | 删除 |
| ----- | ----- | ----- | ----- | ----- | ----- |
| 栈 | - | O(1) | O(n) | - | O(1) |
| 队列 | - | O(1) | O(n) | - | O(1) |



## Array 实战题目

- [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)（腾讯、百度、字节跳动在近半年内面试常考）
- [移动零](https://leetcode-cn.com/problems/move-zeroes/)（华为、字节跳动在近半年内面试常考）
- [爬楼梯](https://leetcode.com/problems/climbing-stairs/)（阿里巴巴、腾讯、字节跳动在半年内面试常考）
- [三数之和](https://leetcode-cn.com/problems/3sum/)（国内、国际大厂历年面试高频老题）

- [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)（腾讯、百度、字节跳动在近半年内面试常考）
- [移动零](https://leetcode-cn.com/problems/move-zeroes/)（华为、字节跳动在近半年内面试常考）
- [爬楼梯](https://leetcode.com/problems/climbing-stairs/)（阿里巴巴、腾讯、字节跳动在半年内面试常考）
- [三数之和](https://leetcode-cn.com/problems/3sum/)（国内、国际大厂历年面试高频老题）

- [两数之和](https://leetcode-cn.com/problems/two-sum/)（近半年内，字节跳动在面试中考查此题达到 152 次）
- [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)（腾讯、百度、字节跳动在近半年内面试常考）
- [移动零](https://leetcode-cn.com/problems/move-zeroes/)（华为、字节跳动在近半年内面试常考）
- [爬楼梯](https://leetcode.com/problems/climbing-stairs/)（阿里巴巴、腾讯、字节跳动在半年内面试常考）
- [三数之和](https://leetcode-cn.com/problems/3sum/)（国内、国际大厂历年面试高频老题）

## Linked List 实战题目

- [反转链表](https://leetcode.com/problems/reverse-linked-list/)（字节跳动、亚马逊在半年内面试常考）
- [两两交换链表中的节点](https://leetcode.com/problems/swap-nodes-in-pairs)（阿里巴巴、字节跳动在半年内面试常考）
- [环形链表](https://leetcode.com/problems/linked-list-cycle)（阿里巴巴、字节跳动、腾讯在半年内面试常考）
- [环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii)
- [K 个一组翻转链表](https://leetcode.com/problems/reverse-nodes-in-k-group/)（字节跳动、猿辅导在半年内面试常考）

## 参考链接

- [Java 源码分析（ArrayList）](http://developer.classpath.org/doc/java/util/ArrayList-source.html)
- [Linked List 的标准实现代码](http://www.geeksforgeeks.org/implementing-a-linked-list-in-java-using-class/)
- [Linked List 示例代码](http://www.cs.cmu.edu/~adamchik/15-121/lectures/Linked Lists/code/LinkedList.java)
- [Java 源码分析（LinkedList）](http://developer.classpath.org/doc/java/util/LinkedList-source.html)
- LRU Cache - Linked list：[ LRU 缓存机制](http://leetcode-cn.com/problems/lru-cache)
- Redis - Skip List：[跳跃表](http://redisbook.readthedocs.io/en/latest/internal-datastruct/skiplist.html)、[为啥 Redis 使用跳表（Skip List）而不是使用 Red-Black？](http://www.zhihu.com/question/20202931)

## 栈参考链接

- [Java 的 PriorityQueue 文档](http://docs.oracle.com/javase/10/docs/api/java/util/PriorityQueue.html)
- [Java 的 Stack 源码](http://developer.classpath.org/doc/java/util/Stack-source.html)
- [Java 的 Queue 源码](http://fuseyism.com/classpath/doc/java/util/Queue-source.html)
- [Python 的 heapq](http://docs.python.org/2/library/heapq.html)
- [高性能的 container 库](http://docs.python.org/2/library/collections.html)



## 栈和队列实战题目

- [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)（亚马逊、JPMorgan 在半年内面试常考）
- [最小栈](https://leetcode-cn.com/problems/min-stack/)（亚马逊在半年内面试常考）

- [柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram)（亚马逊、微软、字节跳动在半年内面试中考过）
- [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum)（亚马逊在半年内面试常考）

- 用 add first 或 add last 这套新的 API 改写 Deque 的代码
- 分析 Queue 和 Priority Queue 的源码
- [设计循环双端队列](https://leetcode.com/problems/design-circular-deque)（Facebook 在 1 年内面试中考过）
- [接雨水](https://leetcode.com/problems/trapping-rain-water/)（亚马逊、字节跳动、高盛集团、Facebook 在半年内面试常考）

说明：改写代码和分析源码这两项作业，同学们需要在第 1 周的学习总结中完成。如果不熟悉 Java 语言，这两项作业可选做。