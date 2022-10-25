/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // o(n*2)
func reversePrint(head *ListNode) []int {
	res := []int{}
	var dfs func(cur *ListNode)
	dfs = func(cur *ListNode) {
		if cur == nil {
			return
		}
		dfs(cur.Next)
		res = append(res, cur.Val)
	}
	dfs(head)
	return res
}
// o(n*1.5)
func reversePrint1(head *ListNode) []int {
	var ret []int
	op := head
	for op != nil {
		ret = append(ret, op.Val)
		op = op.Next
	}
	i := 0
	j := len(ret) - 1
	for i < j {
		tmp := ret[i]
		ret[i] = ret[j]
		ret[j] = tmp
		i++
		j--
	}
	return ret
}