class Solution:
    def containsCycle(self, grid):
        rows, cols = len(grid), len(grid[0])
        visited = [[False] * cols for _ in range(rows)]

        def dfs(r, c, pr, pc):
            if visited[r][c]:
                return True

            visited[r][c] = True

            for dr, dc in [(1,0), (-1,0), (0,1), (0,-1)]:
                nr, nc = r + dr, c + dc

                if 0 <= nr < rows and 0 <= nc < cols:
                    if grid[nr][nc] != grid[r][c]:
                        continue

                    if (nr, nc) == (pr, pc):
                        continue

                    if dfs(nr, nc, r, c):
                        return True

            return False

        for i in range(rows):
            for j in range(cols):
                if not visited[i][j]:
                    if dfs(i, j, -1, -1):
                        return True

        return False