class Solution:
    def mirrorDistance(self, n: int) -> int:
        def rever(a):
            if a // 10 == 0:
                return a
            return int(str(a)[::-1])
        return abs(rever(n)-n)