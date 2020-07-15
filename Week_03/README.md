## 学习笔记
- 这周工作有点忙，详细下周补上
### 递归
- 递归就是通过函数体进行的循环
- 学习要点：不人肉递归、找最近重复性、数学归纳法思维
- 递归模板
```py
def recursion(level, param1, param2, ...): 
    # recursion terminator（递归终止符）
    if level > MAX_LEVEL: 
	   process_result 
	   return 

    # process logic in current level（当前层的处理逻辑）
    process(level, data...) 

    # drill down（下钻）
    self.recursion(level + 1, p1, ...) 

    # reverse the current level status if needed（清理当前层状态）
```
### 分治
- 关键点：分拆子问题、合并结果
- 分治模板
```py
# Python
def divide_conquer(problem, param1, param2, ...): 
  # recursion terminator （递归终止符）
  if problem is None: 
	print_result 
	return 

  # prepare data （分治数据准备）
  data = prepare_data(problem) 
  subproblems = split_problem(problem, data) 

  # conquer subproblems (攻克子问题)
  subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
  subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
  subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
  …

  # process and generate the final result （处理并合并结果）
  result = process_result(subresult1, subresult2, subresult3, …)
	
  # revert the current level states（清理当前层状态）
```
### 回溯
- 回溯就是尝试分布取解决问题，当异常时会回到上一步或几步取继续尝试其它分支
- 典型问题：八皇后
- 分治和回溯就是特殊的递归，较为复杂的递归，本质就是寻找重复性，最优重复性就是动态规划
