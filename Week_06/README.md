#  学习笔记
## 递归模板

```py
# Python
def recursion(level, param1, param2, ...): 
    # recursion terminator (递归终止符)
    if level > MAX_LEVEL: 
	   process_result 
	   return 
    # process logic in current level （当前级别的处理逻辑）
    process(level, data...) 
    # drill down （下钻）
    self.recursion(level + 1, p1, ...) 
    # reverse the current level status if needed  （根据需要反转当前级别状态）
```
## 分治模板
```py
# Python
def divide_conquer(problem, param1, param2, ...): 
  # recursion terminator  (递归终止符)
  if problem is None: 
	print_result 
	return 

  # prepare data （准备数据，分拆）
  data = prepare_data(problem) 
  subproblems = split_problem(problem, data) 

  # conquer subproblems （攻克子问题）
  subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
  subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
  subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
  …

  # process and generate the final result （处理并产生最终结果）
  result = process_result(subresult1, subresult2, subresult3, …)
	
  # revert the current level states（恢复当前级别状态）
```
## 动态规划
- 动态规划 和 递归或者分治 没有根本上的区别（关键看有无最优子结构）
- 动态规划 和 递归或者分治 共性是找到重复子问题
- 动态规划 和 递归或者分治 差异是：DP有最优子结构，中途可以淘汰次优解
- 本质就是分治+最优子结构
## 解题感触
- 人肉递归很累、低效
- 找到最近最简方法，将其拆解成可重复解决的问题
- 数学归纳法思维（抵制人肉递归的诱惑）
