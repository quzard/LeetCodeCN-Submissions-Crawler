/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
    if root == nil {
        return true
    }
    if root.Val <= lower || root.Val >= upper {
        return false
    }
    return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func isValidBST1(root *TreeNode) bool {
	var pre *TreeNode
	flag := true

    var dfs func(n *TreeNode)
	dfs = func(node *TreeNode) {
        if flag == false {
            return
        }
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre == nil {
			pre = node
		} else {
			if node.Val <= pre.Val {
				flag = false
			} else {
				pre = node
			}
		}
		dfs(node.Right)
	}
	dfs(root)
	return flag
}