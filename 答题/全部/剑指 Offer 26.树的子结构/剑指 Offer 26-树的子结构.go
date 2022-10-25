/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil && B == nil {
        return true
    }
    if A == nil || B == nil {
        return false
    }
    var dfs func(A *TreeNode, B *TreeNode) bool
    dfs = func(curA *TreeNode, curB *TreeNode) bool {
        if curA == nil && curB == nil {
            return true
        }
        if curB == nil {
            return true
        }
        if curA == nil {
            return false
        }
        if curA.Val == curB.Val {
            if dfs(curA.Left, curB.Left) && dfs(curA.Right, curB.Right) {
                return true
            }
        }
        if dfs(curA.Left, B) {
            return true
        }
        if dfs(curA.Right, B) {
            return true
        }
        return false
    }
    return dfs(A, B)
}

func isSubStructure1(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil {
        return false
    }
    // A or A.Left or A.Right
    if preOrderContain(A, B) {
        return true
    }

    return isSubStructure((*A).Left, B) || isSubStructure((*A).Right, B)
}

func preOrderContain(a *TreeNode, b *TreeNode) bool {
    if b == nil {
        return true
    }
    if a == nil {
        return false
    }

    if (*a).Val != (*b).Val {
        return false
    }

    bl := preOrderContain((*a).Left, (*b).Left)
    br := preOrderContain((*a).Right, (*b).Right)

    return bl && br
}