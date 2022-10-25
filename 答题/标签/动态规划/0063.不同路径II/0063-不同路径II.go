func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	dp := [2][]int{}
	for i := 0; i < len(obstacleGrid); i++ {
		dp[1] = make([]int, len(obstacleGrid[0]))
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 {
				if j == 0 {
					if obstacleGrid[i][j] == 1 {
						dp[1][j] = 0
					} else {
						dp[1][j] = 1
					}
				} else {
					if obstacleGrid[i][j] == 1 {
						dp[1][j] = 0
					} else {
						dp[1][j] = dp[1][j-1]
					}
				}
			} else {
				if j == 0 {
					if obstacleGrid[i][j] == 1 {
						dp[1][j] = 0
					} else {
						dp[1][j] = dp[0][j]
					}
				} else {
					if obstacleGrid[i][j] == 1 {
						dp[1][j] = 0
					} else {
						dp[1][j] = dp[0][j] + dp[1][j-1]
					}
				}
			}
		}
		dp[0] = dp[1]
	}
	return dp[0][len(dp[0])-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	path := make([][]int, n+1)
	for i := range path {
		path[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				path[i][j] = 0
			} else if i == 1 && j == 1 {
				path[i][j] = 1
			} else {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}

		}
	}
	return path[n][m]
}