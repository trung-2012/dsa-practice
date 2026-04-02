class Solution:
    def hammingWeight(self, n: int) -> int:
        dem = 0
        while n > 0 :
            if n % 2 == 1 :
                dem += 1
            n //= 2    
        return dem            