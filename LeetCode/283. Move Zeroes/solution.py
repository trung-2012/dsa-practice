class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        m = 0
        for i in range(len(nums)):
            if nums[i]!=0:
                if m !=i:
                    nums[m] = nums[i]
                    nums[i] = 0
                m += 1
        return nums
        