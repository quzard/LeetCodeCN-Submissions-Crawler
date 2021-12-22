/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res [][]int

func dfs(root *TreeNode, high int){
    if root == nil{
        return
    }
    if len(res) < high + 1{
        res = append(res, []int{})
    }
    if high %2 == 0{
        res[high] = append(res[high], root.Val)
        dfs(root.Left, high + 1)
        dfs(root.Right, high + 1)
    }else{
        res[high] = append(append([]int{}, root.Val),res[high]...)
        dfs(root.Left, high + 1)
        dfs(root.Right, high + 1)
    }
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    res = [][]int{}
    dfs(root, 0)
    return res
}