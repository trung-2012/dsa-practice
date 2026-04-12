class Solution:
    def numberOfChild(self, n: int, k: int) -> int:
        cycle = 2 * (n - 1)
        pos = k % cycle

        if pos <= n - 1:
            return pos
        else:
            return cycle - pos