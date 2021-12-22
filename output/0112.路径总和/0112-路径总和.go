/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum1(root *TreeNode, targetSum int) bool {
    if root == nil{
        return false
    }
    res := false
    var dfs func(cur *TreeNode, sum int)
    dfs = func(cur *TreeNode, sum int){
        if res || cur == nil{
            return
        }
        sum += cur.Val
        if sum == targetSum && cur.Right == nil && cur.Left == nil{
            res = true
            return 
        }
        dfs(cur.Left, sum)
        dfs(cur.Right, sum)
    }
    dfs(root, 0)
    return res
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }

    return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}