class Solution:
    def hasValidPath(self, grid: List[List[int]]) -> bool:
        m, n = len(grid), len(grid[0])

        directions = {
            1: [(0, -1), (0, 1)],
            2: [(-1, 0), (1, 0)],
            3: [(0, -1), (1, 0)],
            4: [(0, 1), (1, 0)],
            5: [(0, -1), (-1, 0)],
            6: [(0, 1), (-1, 0)],
        }

        opposite = {
            (0, -1): (0, 1),
            (0, 1): (0, -1),
            (-1, 0): (1, 0),
            (1, 0): (-1, 0),
        }

        visited = set()
        q = deque([(0, 0)])
        visited.add((0, 0))

        while q:
            x, y = q.popleft()

            if (x, y) == (m - 1, n - 1):
                return True

            for dx, dy in directions[grid[x][y]]:
                nx, ny = x + dx, y + dy

                if 0 <= nx < m and 0 <= ny < n:
                    if (nx, ny) in visited:
                        continue
                    if opposite[(dx, dy)] in directions[grid[nx][ny]]:
                        visited.add((nx, ny))
                        q.append((nx, ny))

        return False