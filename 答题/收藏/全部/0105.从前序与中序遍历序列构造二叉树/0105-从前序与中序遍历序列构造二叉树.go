/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    i := 0
    for ; i < len(inorder) && inorder[i] != preorder[0]; i++ {}

    mid := &TreeNode{Val: preorder[0]}

    mid.Left = buildTree(preorder[1:1+i], inorder[:i])
    mid.Right = buildTree(preorder[1+i:], inorder[i+1:])
    return mid
}