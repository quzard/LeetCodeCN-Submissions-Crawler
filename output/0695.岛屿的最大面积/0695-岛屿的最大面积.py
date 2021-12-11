class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        def dfs(i,j):
            if grid[i][j] != 1:
                return 0
            grid[i][j] = 0
            a = b = c = d = 0
            if i>0:
                a = dfs(i-1,j)
            if j>0:
                b = dfs(i,j-1)
            if i<m-1:
                c = dfs(i+1,j)
            if j < n-1:
                d = dfs(i,j+1)
            return a+b+c+d+1
        res = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j]:
                    res  = max(res,dfs(i,j))
        return res