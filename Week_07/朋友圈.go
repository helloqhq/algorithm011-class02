package Week_07

// 使用并查集
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	length := len(M)
	uFind := NewUnionFind(length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ { // 对称矩阵 不需要重复计算
			if M[i][j] == 1 {
				uFind.Union(i, j)
			}
		}
	}
	return uFind.count
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {//还可以小树接大树，避免树太高，尽量保持树平衡
		u.parent[pi] = pj
		u.count--
	}
}

// 查找当前节点所在集合的根节点
// 路径压缩后 查询为O(1)
func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
