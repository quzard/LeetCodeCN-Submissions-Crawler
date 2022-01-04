/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(preorder) != len(inorder){
        return nil
    }
    if len(preorder) == 1{
        return &TreeNode{Val: preorder[0]}
    }
    head := &TreeNode{Val: preorder[0]}
    i:=0
    for i = 0; i< len(inorder);i++{
        if inorder[i] == preorder[0]{
            break
        }
    }
    head.Left = buildTree(preorder[1:1+i], inorder[:i])
    head.Right = buildTree(preorder[i+1:], inorder[i+1:])
   
    return head
}