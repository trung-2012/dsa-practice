class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        res = []
        min_diff = float('inf')

        for i in range(1, len(arr)):
            M = arr[i] - arr[i-1]
            if M < min_diff:
                min_diff = M
                res = [[arr[i-1], arr[i]]]
            elif M == min_diff:
                res.append([arr[i-1], arr[i]])
        return res