class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        count = {'L': 0, 'R': 0, '_': 0}

        for ch in moves:
            count[ch] += 1
        if count['L'] > count['R']:
            return count['L'] - count['R'] + count['_']
        return count['R'] - count['L'] + count['_']