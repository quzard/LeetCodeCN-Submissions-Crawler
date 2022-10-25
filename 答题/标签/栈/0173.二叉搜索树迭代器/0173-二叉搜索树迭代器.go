/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    list []int
    cur  int
}


func Constructor(root *TreeNode) BSTIterator {
    list := []int{}
    var dfs func(cur *TreeNode)
    dfs = func(cur *TreeNode){
        if cur == nil{
            return
        }
        dfs(cur.Left)
        list = append(list, cur.Val)
        dfs(cur.Right)
    }
    dfs(root)
    return BSTIterator{
        list: list,
        cur: 0,
    }
}


func (this *BSTIterator) Next() int {
    if this.HasNext(){
        res := this.list[this.cur]
        this.cur++
        return res
    }
    return 0
}


func (this *BSTIterator) HasNext() bool {
    return this.cur < len(this.list)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */