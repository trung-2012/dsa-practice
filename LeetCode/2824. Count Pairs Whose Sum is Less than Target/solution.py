class Solution:
    def countPairs(self, nums: List[int], target: int) -> int:
        count = 0
        nums.sort()
        left = 0
        right = len(nums)

        while left < right:
            if nums[left] + nums[right-1] < target:
                count += right - left - 1
                left += 1
            else:
                right -= 1
        return count
