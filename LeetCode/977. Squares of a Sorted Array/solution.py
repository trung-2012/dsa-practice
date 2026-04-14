class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        res = []
        for x in nums:
            res.append(x**2)
        res.sort()
        return res