package Week_02

import "strings"

func removeOuterParentheses(s string) string {
	sSlice := []byte(s)
	ans := strings.Builder{}
	stack := make([]int, 0)

	start := 0    // 初始化原语的起始位置
	end := 0      // 初始化原语的结束位置
	flag := false // 标志每个原语

	for k, v := range sSlice {
		if v == '(' { // 遇到左括号，入栈
			stack = append(stack, '(')
			if !flag { // 遇到的第一个左括号，是原语的开始位置，记录下原语开始位置
				start = k
				flag = true
			}
		}

		if v == ')' { // 遇到右括号，出栈
			stack = stack[1:]
			if len(stack) == 0 { // 当栈空的时候，找到了一个完整的原语
				end = k                          // 记录下结束位置
				ans.Write(sSlice[start+1 : end]) // 去掉原语的最外层括号，并追加到答案中
				flag = false                     // 置标志为false，往后接着找下一个原语
				start = end                      // 往后找，再次初始化原语开始位置
			}
		}
	}

	return ans.String()
}

