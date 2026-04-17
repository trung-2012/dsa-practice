class Solution:
    def deleteAndEarn(self, nums: List[int]) -> int:
        if not nums:
            return 0
        freq = [0]*(max(nums)+1)
        for n in nums:
            freq[n] += n
        
        prev = 0
        cur = 0

        for val in freq:
            temp = cur
            cur = max(cur, prev + val)
            prev = temp
        return cur