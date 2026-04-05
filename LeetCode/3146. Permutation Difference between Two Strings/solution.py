class Solution:
    def findPermutationDifference(self, s: str, t: str) -> int:
        post = {}
        res = 0

        for i in range(len(t)):
            post[t[i]] = i
        for i in range(len(s)):
            res += abs(i - post[s[i]])
        return res