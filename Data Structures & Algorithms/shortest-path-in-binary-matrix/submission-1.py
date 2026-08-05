from collections import deque

class Solution:
    def shortestPathBinaryMatrix(self, grid):
        rows, cols = len(grid), len(grid[0])

        if grid[0][0] == 1 or grid[rows - 1][cols - 1] == 1:
            return -1

        directions = [
            (0, 1), (0, -1),
            (1, 0), (-1, 0),
            (1, 1), (1, -1),
            (-1, 1), (-1, -1)
        ]

        visited = {(0, 0)}
        queue = deque()
        queue.append((0, 0))

        distance = 1

        while queue:
            for _ in range(len(queue)):
                r, c = queue.popleft()

                if r == rows - 1 and c == cols - 1:
                    return distance

                for dr, dc in directions:
                    nr, nc = r + dr, c + dc

                    if (
                        0 <= nr < rows
                        and 0 <= nc < cols
                        and grid[nr][nc] == 0
                        and (nr, nc) not in visited
                    ):
                        visited.add((nr, nc))
                        queue.append((nr, nc))

            distance += 1

        return -1