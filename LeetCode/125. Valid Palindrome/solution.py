class Solution:
    def isPalindrome(self, s: str) -> bool:
        filtered = []
        
        for c in s:
            if c.isalnum():
                filtered.append(c.lower())
        
        filtered = "".join(filtered)
        return filtered == filtered[::-1]