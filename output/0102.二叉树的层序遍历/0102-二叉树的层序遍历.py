# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        ans = deque([root])
        res = []
        def bfs():
            while ans:
                temp = []
                for _ in range(len(ans)):
                    now = ans.popleft()
                    temp.append(now.val)
                    if now.left:
                        ans.append(now.left)
                    if now.right:
                        ans.append(now.right)
                res.append(temp)
        bfs()
        return res