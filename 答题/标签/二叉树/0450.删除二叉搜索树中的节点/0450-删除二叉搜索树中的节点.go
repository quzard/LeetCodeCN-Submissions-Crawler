/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteNode(root *TreeNode, key int) *TreeNode {
	var dfs func(cur *TreeNode) *TreeNode
	dfs = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}
		var node *TreeNode
		if cur.Val == key {
			if cur.Right != nil {
				node = cur.Right
				if cur.Left != nil {
					temp := node.Left
					node.Left = cur.Left
					cur = node.Left
					for cur.Right != nil {
						cur = cur.Right
					}
					cur.Right = temp
				}
				return node
			} else {
				return cur.Left
			}
		} else if cur.Val < key {
			cur.Right = dfs(cur.Right)
		} else {
			cur.Left = dfs(cur.Left)
		}

		return cur
	}
	return dfs(root)
}

func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		// 只有一个非空子节点，那么让孩子接替自己的位置
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 有两个子节点，寻找左边或者右边接替自己，这里直接进行数据的交换就行
		// 然后递归删除旧数据
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}