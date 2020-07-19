# 学习笔记
## 深度优先
### 递归模板
```py
#Python
visited = set() 

def dfs(node, visited):
    if node in visited: # terminator
    	# already visited 
    	return 

	visited.add(node) 

	# process current node here. 
	...
	for next_node in node.children(): 
		if next_node not in visited: 
			dfs(next_node, visited)
```
### 非递归模板（借助栈实现）
```py
#Python
def DFS(self, tree): 

	if tree.root is None: 
		return [] 

	visited, stack = [], [tree.root]

	while stack: 
		node = stack.pop() 
		visited.add(node)

		process (node) 
		nodes = generate_related_nodes(node) 
		stack.push(nodes) 

	# other processing work 
	...
```
## 广度优先
### 模板（借助队列实现）
```py
# Python
def BFS(graph, start, end):
    visited = set()
	queue = [] 
	queue.append([start]) 
	while queue: 
		node = queue.pop() 
		visited.add(node)
		process(node) 
		nodes = generate_related_nodes(node) 
		queue.push(nodes)
	# other processing work 
	...
```
## 贪心算法
- 每一步都取最优解，不一定能得到全局最优
- 贪心和动归的不同是对子问题的处理方式，贪心每一做出选择后不能回退，动归会保存以前的结果，并根据以前结果对当前进行选择，有回退功能
- 贪心：当下做局部最优判断
- 回溯：能够回退
- 动态规划：最优判断+回退
## 二分查找
### 前提
- 值有序
- 存在上下界
- 能通过索引访问
### 模板
```py
# Python
left, right = 0, len(array) - 1 
while left <= right: 
	  mid = (left + right) / 2 
	  if array[mid] == target: 
		    # find the target!! 
		    break or return result 
	  elif array[mid] < target: 
		    left = mid + 1 
	  else: 
		    right = mid - 1
```
