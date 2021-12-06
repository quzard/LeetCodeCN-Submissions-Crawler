class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(len(nums) - 1, 0 , -1):
            if nums[i] == nums[i-1]:
                return nums[i]