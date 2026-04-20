class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        maxI = 0

        for i in range(len(colors)-1, -1, -1):
            if colors[i] != colors[0]:
                maxI = max(maxI, i)
                break
        
        for i in range(len(colors)):
            if colors[i] != colors[len(colors)-1]:
                maxI = max(maxI, len(colors)-1-i)
                break
        return maxI