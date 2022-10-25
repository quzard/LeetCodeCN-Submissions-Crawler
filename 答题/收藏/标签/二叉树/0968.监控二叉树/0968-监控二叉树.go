/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	// 采用后后序遍历从最底下的节点开始，如果该节点没被观测到且存在父节点，就给父节点装摄像头。
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 1
		}

		if node.Left == nil && node.Right == nil {
			return 0
		}
		// 0 表示叶节点
		// 1 表示nil
		// 2 表示装了摄像头

		l := dfs(node.Left)
		r := dfs(node.Right)

        //存在没被观测的子节点, 当前节点装摄像头
		if l == 0 || r == 0 {
			ans++
			return 2
		}

        //子节点装了摄像头, 父节点不用关心当前节点
		if l == 2 || r == 2 {
			return 1
		}

        //子节点都不需要当前关心, 则当前节点可假设为叶节点
		return 0
	}
	if dfs(root) == 0 {
		ans++
	}

	return ans
}