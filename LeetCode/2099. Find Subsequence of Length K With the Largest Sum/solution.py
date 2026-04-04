class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        s = sorted(nums, reverse = True)[:k]
        res = []
        for x in nums:
            if x in s:
                res.append(x)
                s.remove(x)

        return res