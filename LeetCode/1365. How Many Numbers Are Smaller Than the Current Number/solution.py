class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        sorted_nums = sorted(nums)

        pos = {}

        for i, v in enumerate(sorted_nums):
            if v not in pos:
                pos[v] = i
        return [pos[x] for x in nums]