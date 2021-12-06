# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# class Solution:
#     def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
class Solution:
    # res = []
    def pathSum(self, root: TreeNode, targetSum: int):
        # self.res = []
        res = []
        path = []
        # @functools.lru_cache(None)
        def dfs(p, val):
            if p.val == val and not p.left and not p.right:
                res.append(path[:] + [p.val])
            # elif val < p.val :
            #     return
            else:
                path.append(p.val)
                if p.left:
                    dfs(p.left, val - p.val)
                if p.right:
                    dfs(p.right, val - p.val)
                path.pop()
        if root:
            dfs(p=root, val=targetSum)
        return res