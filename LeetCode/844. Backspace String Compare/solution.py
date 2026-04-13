class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        def build(x):
            stack = []

            for i in x:
                if i != "#":
                    stack.append(i)
                elif stack:
                    stack.pop()
            return "".join(stack)
        return build(s) == build(t)