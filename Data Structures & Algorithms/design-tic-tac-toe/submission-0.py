class TicTacToe:

    def __init__(self, n: int):
        self.rows = [0] * n
        self.cols = [0] * n
        self.diagonal = 0
        self.antiDiagonal = 0

    def move(self, row: int, col: int, player: int) -> int:
        currentPlayer = 1 if player == 1 else -1

        # update currentPlayer in rows and cols arrays
        self.rows[row] += currentPlayer
        self.cols[col] += currentPlayer

        # update diagonal
        if row == col:
            self.diagonal += currentPlayer

        # update anti diagonal
        if col == (len(self.cols) - row - 1):
            self.antiDiagonal += currentPlayer

        n = len(self.rows)

        # check if the current player wins
        if (abs(self.rows[row]) == n or
            abs(self.cols[col]) == n or
            abs(self.diagonal) == n or
            abs(self.antiDiagonal) == n):
            return player

        # No one wins
        return 0