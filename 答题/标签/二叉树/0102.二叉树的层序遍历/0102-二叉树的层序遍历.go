/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, index int)
	dfs = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if index > len(res) {
			res = append(res, []int{root.Val})
		} else {
			res[index-1] = append(res[index-1], root.Val)
		}
		dfs(root.Left, index+1)
		dfs(root.Right, index+1)
	}
	dfs(root, 1)
	return res
}