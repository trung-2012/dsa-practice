class Solution:
    def isHappy(self, n: int) -> bool:
        def check(m):
            res = 0
            while m > 0 :
                a = m % 10
                m //= 10
                res += a**2
            return res
        seen = set()

        while n != 1 and n not in seen:
            seen.add(n)
            n = check(n)
        return n == 1