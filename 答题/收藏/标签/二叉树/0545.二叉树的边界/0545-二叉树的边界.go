/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    if nil == root {
		return []int{}
	}
	res := []int{root.Val}

    var dfs func(cur *TreeNode, l, r bool)
    dfs = func(cur *TreeNode, l, r bool){
        if cur == nil{
            return
        }
        if cur.Left == nil && cur.Right == nil{
            res = append(res, cur.Val)
            return
        }
        if l {
            res = append(res, cur.Val)
        }

        dfs(cur.Left, l , r && nil == cur.Right)
		dfs(cur.Right, l && nil == cur.Left, r )

        if r {
            res = append(res, cur.Val)
        }
        
    }
    dfs(root.Left, true, false)
    dfs(root.Right, false, true)


    return res
}
