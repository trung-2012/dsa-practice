class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 3: return n

        prev1 = 3
        prev2 = 2
        temp = 0

        for _ in range(3, n):
            temp = prev1 + prev2
            prev2 = prev1
            prev1 = temp
        return temp