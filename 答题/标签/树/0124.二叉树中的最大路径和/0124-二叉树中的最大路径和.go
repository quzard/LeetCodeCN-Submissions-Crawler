/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum1(root *TreeNode) int {
    res := math.MinInt32
    m := map[*TreeNode]int {}
    var dfs func(cur *TreeNode)
    dfs = func(cur *TreeNode) {
        if cur == nil {
            return
        }
        dfs(cur.Left)
        dfs(cur.Right)
        num1, num2 := math.MinInt32, math.MinInt32
        if cur.Left != nil {
            num1 = m[cur.Left]
        }
        if cur.Right != nil {
            num2 = m[cur.Right]
        }
        num := cur.Val
        m[cur] = max(max(num1, num2) + num, num)
        res = max(res, max(max(max(num1 + num, num2 + num), num), num1 + num2 + num))
    }
    dfs(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}



func maxPathSum(root *TreeNode) int {
    res := math.MinInt32
    var dfs func(root *TreeNode)int
    dfs = func(root *TreeNode)int{
        if root == nil{
            return 0
        }
        left := max(dfs(root.Left), 0)
        right := max(dfs(root.Right), 0)
        res = max(res, left + right + root.Val)
        return root.Val+max(left, right)
    }
    res = max(res, dfs(root))
    return res
}