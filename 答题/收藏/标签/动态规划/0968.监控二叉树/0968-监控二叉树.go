/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	// ���ú�������������µĽڵ㿪ʼ������ýڵ�û���۲⵽�Ҵ��ڸ��ڵ㣬�͸����ڵ�װ����ͷ��
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 1
		}

		if node.Left == nil && node.Right == nil {
			return 0
		}
		// 0 ��ʾҶ�ڵ�
		// 1 ��ʾnil
		// 2 ��ʾװ������ͷ

		l := dfs(node.Left)
		r := dfs(node.Right)

        //����û���۲���ӽڵ�, ��ǰ�ڵ�װ����ͷ
		if l == 0 || r == 0 {
			ans++
			return 2
		}

        //�ӽڵ�װ������ͷ, ���ڵ㲻�ù��ĵ�ǰ�ڵ�
		if l == 2 || r == 2 {
			return 1
		}

        //�ӽڵ㶼����Ҫ��ǰ����, ��ǰ�ڵ�ɼ���ΪҶ�ڵ�
		return 0
	}
	if dfs(root) == 0 {
		ans++
	}

	return ans
}