# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        res = []
        q = deque([root])
        while q:
            now = q.popleft()
            res.append(now.val)
            if now.left:
                q.append(now.left)
            if now.right:
                q.append(now.right)
        return res