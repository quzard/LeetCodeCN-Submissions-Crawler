/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    // inorder   中序 左中右
    // postorder 后序 左右中
    if len(postorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: postorder[len(postorder) - 1]}
    i := 0
    for i < len(inorder) {
        if inorder[i] == root.Val {
            break
        }
        i++
    }
    root.Left = buildTree(inorder[:i], postorder[:i])
    root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
    return root
}