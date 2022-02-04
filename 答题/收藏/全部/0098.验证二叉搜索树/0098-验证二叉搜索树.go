/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST1(root *TreeNode) bool {
    res := true
    var dfs func(root*TreeNode, l, r int) // dir true right. false  left
    dfs = func(root*TreeNode, l, r int){
        if res == false{
            return
        }
        if l != math.MinInt64{
            if root.Val <= l{
                res = false
                return
            }
        }
        if r != math.MaxInt64{
            if root.Val >= r{
                res = false
                return
            }
        }
        if root.Left != nil{
            if root.Left.Val < root.Val{
                dfs(root.Left, l, min(root.Val, r))
            }else{
                res = false
                return 
            }
        }
        if root.Right!=nil{
            if root.Right.Val > root.Val{
                dfs(root.Right, min(root.Val, l), r)
            }else{
                res = false
                return 
            }
        }
    }


    if root.Left != nil{
        if root.Left.Val < root.Val{
            dfs(root.Left, math.MinInt64, root.Val)
        }else{
            return false
        }
    }
    if root.Right!=nil{
        if root.Right.Val > root.Val{
            dfs(root.Right, root.Val, math.MaxInt64)
        }else{
            return false
        }
    }

    return res
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}






func isValidBST(root *TreeNode) bool {

    var inOrder func(n *TreeNode)
	var pre *TreeNode
	flag := true

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		if pre == nil {
			pre = node
		} else {
			if node.Val <= pre.Val {
				flag = false
			} else {
				pre = node
			}
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return flag

}