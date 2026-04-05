class Solution:
    def minimumRightShifts(self, nums: List[int]) -> int:
        n = len(nums)
        count = 0
        idx = -1

        for i in range(n-1):
            if nums[i] > nums[i+1]:
                count += 1
                idx = i
        if count == 0:
            return 0
        if count > 1:
            return -1
        if nums[-1] > nums[0]:
            return -1
        return n - idx - 1