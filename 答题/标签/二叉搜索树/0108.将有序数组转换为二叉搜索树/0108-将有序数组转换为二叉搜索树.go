/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0{
        return nil
    }
    l := 0
    r := len(nums) - 1
    mid := int(l + (r - l) / 2)
    node := &TreeNode{Val:nums[mid]}
    node.Left = sortedArrayToBST(nums[:mid])
    node.Right = sortedArrayToBST(nums[mid+1:])
    return node
}