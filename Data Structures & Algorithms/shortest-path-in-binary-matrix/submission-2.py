from collections import deque


class Solution:
    def shortestPathBinaryMatrix(self, grid):
        rows = len(grid)
        cols = len(grid[0])

        if grid[0][0] == 1 or grid[rows-1][cols-1] == 1:
            return -1

        directions = [
            (0, 1), (0, -1),
            (1, 0), (-1, 0),
            (1, 1), (1, -1),
            (-1, 1), (-1, -1),
        ]

        visited = set()
        visited.add((0, 0))

        queue = deque()
        queue.append((0, 0))

        distance = 1

        while len(queue) > 0:
            size = len(queue)

            for i in range(size):
                current = queue.popleft()

                r = current[0]
                c = current[1]

                if r == rows-1 and c == cols-1:
                    return distance

                for direction in directions:
                    dr = direction[0]
                    dc = direction[1]

                    nr = r + dr
                    nc = c + dc

                    if (
                        nr >= 0 and nr < rows and
                        nc >= 0 and nc < cols and
                        grid[nr][nc] == 0 and
                        (nr, nc) not in visited
                    ):
                        visited.add((nr, nc))
                        queue.append((nr, nc))

            distance += 1

        return -1