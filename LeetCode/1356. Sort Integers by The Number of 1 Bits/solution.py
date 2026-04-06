class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        temp = []
        for x in arr:
            count = bin(x).count('1')
            temp.append((count, x))
        temp.sort()
        res = []
        for a, b in temp:
            res.append(b)
        return res