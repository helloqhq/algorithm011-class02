package Week_04

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gLen,sLen:=len(g),len(s)
	i, j := 0, 0
	for i < gLen && j < sLen  {//贪心算法：由小到大遍历，优先用小饼干满足胃口小的熊孩子
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}