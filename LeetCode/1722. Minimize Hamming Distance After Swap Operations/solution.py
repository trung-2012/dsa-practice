class Solution:
    def minimumHammingDistance(self, source: List[int], target: List[int], allowedSwaps: List[List[int]]) -> int:
        n = len(source)
        parent = list(range(n))

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x, y):
            parent[find(x)] = find(y)

        for a, b in allowedSwaps:
            union(a, b)

        groups = defaultdict(list)
        for i in range(n):
            groups[find(i)].append(i)

        res = 0
        for group in groups.values():
            count = {}
            for i in group:
                count[source[i]] = count.get(source[i], 0) + 1
            for i in group:
                if count.get(target[i], 0) > 0:
                    count[target[i]] -= 1
                else:
                    res += 1

        return res