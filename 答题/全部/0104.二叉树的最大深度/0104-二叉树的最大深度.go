/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth1(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode, index int)
    dfs = func(root *TreeNode, index int){
        if root != nil{
            res = max(res , index)
            dfs(root.Left, index + 1)
            dfs(root.Right, index + 1)
        }
    }
    dfs(root, 1)

    return res
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []TreeNode
	queue = append(queue, *root)
	var result int
	for ;len(queue) != 0; {
		size := len(queue)
		result ++
		for ;size > 0; size -- {
			if queue[0].Left != nil {
				queue = append(queue, *queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, *queue[0].Right)
			}
			queue = queue[1:]
		}
	}

	return result
}
func max(a, b int) int{
    if a > b{
        return a
    }else{
        return b
    }
}