class Solution:
    def distance(self, nums: List[int]) -> List[int]:
        groups = defaultdict(list)
        
        for i, num in enumerate(nums):
            groups[num].append(i)
        
        res = [0] * len(nums)
        
        for indices in groups.values():
            prefix = [0] * len(indices)
            prefix[0] = indices[0]
            
            for i in range(1, len(indices)):
                prefix[i] = prefix[i - 1] + indices[i]
            
            total = prefix[-1]
            k = len(indices)
            
            for j, idx in enumerate(indices):
                left = idx * j - (prefix[j - 1] if j > 0 else 0)
                right = (total - prefix[j]) - (k - j - 1) * idx
                res[idx] = left + right
        
        return res