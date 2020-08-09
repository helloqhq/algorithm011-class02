# 学习笔记

## 树

###  二叉搜索树Binary Search Tree
- 二叉搜索树，也称二叉搜索树、有序二叉树（Ordered Binary Tree）、排序二叉树（Sorted Binary Tree），是指一棵空树或者具有下列性质的二叉树：
    - 左子树上所有结点的值均小于它的根结点的值；
    - 右子树上所有结点的值均大于它的根结点的值；
    - 以此类推：左、右子树也分别为二叉查找树。（这就是重复性！）
- 二叉搜索树的中序遍历结果是有序的
- 时间复杂度是数的深度

### 字典树

- 定义：即 Trie 树，又称单词查找树或键树，是一种树形结构。
- 应用场景：典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
- 优点：最大限度地减少无谓的字符串比较，查询效率比哈希表高。
- 基本性质：
    - 结点本身不存完整单词；
    - 从根结点到某一结点，路径上经过的字符连接起来，为该结点对应的字符串；
    - 每个结点的所有子结点路径代表的字符都不相同。
- 代码模板
```py
# Python 
class Trie(object):
  
	def __init__(self): 
		self.root = {} 
		self.end_of_word = "#" 
 
	def insert(self, word): 
		node = self.root 
		for char in word: 
			node = node.setdefault(char, {}) 
		node[self.end_of_word] = self.end_of_word 
 
	def search(self, word): 
		node = self.root 
		for char in word: 
			if char not in node: 
				return False 
			node = node[char] 
		return self.end_of_word in node 
 
	def startsWith(self, prefix): 
		node = self.root 
		for char in prefix: 
			if char not in node: 
				return False 
			node = node[char] 
		return True
```
### 保证性能的关键：保持树的平衡
### AVL树
- 发明者G. M. Adelson-Velsky和Evgenii Landis
- Balance Factor（平衡因子）：是它的左子树的高度减去它的右子树的高度（有时相反）。
- balance factor = {-1, 0, 1}
- 平衡二叉搜索树
- 通过旋转操作来进行平衡（四种）
    - 左旋


  ![image](../image.png)

    - 右旋

  ![image_1](../image_1.png)
    - 左右旋

  ![image_2](../image_2.png)
  ![image_3](../image_3.png)
  ![image_4](../image_4.png)
    - 右左旋

  ![image_5](../image_5.png)
  ![image_6](../image_6.png)
  ![image_7](../image_7.png)
- 不足：结点需要存储额外信息、且调整次数频繁
### 红黑树
- 定义：红黑树是一种近似平衡的二叉搜索树（BinarySearch Tree），它能够确保任何一个结点的左右子树的高度差小于两倍
- 红黑树是满足如下条件的二叉搜索树：
  - 每个结点要么是红色，要么是黑色
  - 根结点是黑色
  - 每个叶结点（NIL结点，空结点）是黑色的
  - 不能有相邻接的两个红色结点
  - 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点
## 并查集
- 适用场景：组团或配对的问题
- 基本操作
    - makeSet(s)：建立一个新的并查集，其中包含 s 个单元素集合。
    - unionSet(x, y)：把元素 x 和元素 y 所在的集合合并，要求 x 和 y 所在的集合不相交，如果相交则不合并。
    - find(x)：找到元素 x 所在的集合的代表，该操作也可以用于判断两个元素是否位于同一个集合，只要将它们各自的代表比较一下就可以了。
- 代码模板
```py
# Python 
def init(p): 
	# for i = 0 .. n: p[i] = i; 
	p = [i for i in range(n)] 
 
def union(self, p, i, j): 
	p1 = self.parent(p, i) 
	p2 = self.parent(p, j) 
	p[p1] = p2 
 
def parent(self, p, i): 
	root = i 
	while p[root] != root: 
		root = p[root] 
	while p[i] != i: # 路径压缩 ?
		x = i; i = p[i]; p[x] = root 
	return root
```
## 回溯法
- 定义：回溯法采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。
- 回溯法通常用最简单的递归方法来实现，在反复重复上述的步骤后可能出现两种情况：
    - 找到一个可能存在的正确的答案
    - 在尝试了所有可能的分步方法后宣告该问题没有答案
- 在最坏的情况下，回溯法会导致一次复杂度为指数时间的计算。
## 双向BFS
## A* search
- 启发式函数： h(n)，它用来评价哪些结点最有希望的是一个我们要找的结点，h(n) 会返回一个非负实数,也可以认为是从结点n的目标结点路径的估计成本。
- 启发式函数是一种告知搜索方向的方法。它提供了一种明智的方法来猜测哪个邻居结点会导向一个目标。
- 代码模板
```py
# Python
def AstarSearch(graph, start, end):
	pq = collections.priority_queue() # 优先级 —> 估价函数
	pq.append([start]) 
	visited.add(start)
	while pq: 
		node = pq.pop() # can we add more intelligence here ?
		visited.add(node)
		process(node) 
		nodes = generate_related_nodes(node) 
   unvisited = [node for node in nodes if node not in visited]
		pq.push(unvisited)
```
