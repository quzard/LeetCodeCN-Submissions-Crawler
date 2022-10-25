/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// ¶þ²æËÑË÷Ê÷
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    found := false
    var res *TreeNode
    var dfs func(cur *TreeNode)
    dfs = func(cur *TreeNode){
        if cur == nil || res != nil{
            return
        }
        if cur == p{
            found = true
            dfs(cur.Right)
        }else{
            if p.Val > cur.Val{
                if found && res == nil{
                    res = cur
                    return
                }
                dfs(cur.Right)
            }else{
                dfs(cur.Left)
                if found && res == nil{
                    res = cur
                }
            }
        }
    }
    dfs(root)
    return res
}

func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
    if root == nil || p == nil {
		return nil
	}

	if root.Val > p.Val {
		node := inorderSuccessor(root.Left, p)
		if node == nil {
			return root
		} else {
			return node
		}
	} else {
		node := inorderSuccessor(root.Right, p)
		return node
	}
}