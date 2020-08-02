package Week_06

func minPathSum(grid [][]int) int {
	for i, v := range grid {
		for j, vv := range v {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				grid[i][j] = grid[i][j-1] + vv
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + vv
			} else {
				if grid[i][j-1] < grid[i-1][j] {
					grid[i][j] = grid[i][j-1] + vv
				} else {
					grid[i][j] = grid[i-1][j] + vv
				}
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
