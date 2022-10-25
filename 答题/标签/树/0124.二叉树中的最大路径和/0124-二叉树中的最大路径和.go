/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    res := math.MinInt
    var dfs func(cur *TreeNode) int
    dfs = func(cur *TreeNode) int {
        if cur == nil {
            return 0
        }
        sum := cur.Val
        left := dfs(cur.Left)
        right := dfs(cur.Right)
        res = max(res, sum, sum + left, sum + right, sum + left + right)
        sum = max(max(sum, sum + left), sum + right)
        return sum
    }
    dfs(root)
    return res
}

func max(nums ...int) int {
    m := nums[0]
    for _, num := range nums {
        if num > m {
            m = num
        }
    }
    return m
}