/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    // preorder 前序 中左右
    // inorder  中序 左中右
    if len(preorder) == 0{
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    i:=0
    for i< len(inorder){
        if inorder[i] == preorder[0]{
            break
        }
        i++
    }
    root.Left = buildTree(preorder[1:1+i], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
   
    return root
}