package maps

const MAX_COL = 9
const MAX_RAW = 9
const SQUARE_SIDE = 3

func IsValidSudoku(board [][]byte) bool {
	if checkRaws(board) && checkCols(board) && checkSquares(board) {
		return true
	}
	return false
}

func checkRaws(board [][]byte) bool {
	var flags map[byte]bool = make(map[byte]bool)

	for i := 0; i < MAX_RAW; i++ {
		for _, element := range board[i] {
			if element == '.' {
				continue
			}
			_, exist := flags[element]
			if exist {
				return false
			}
			flags[element] = true
		}
		flags = make(map[byte]bool)
	}

	return true
}

func checkCols(board [][]byte) bool {
	var flags map[byte]bool = make(map[byte]bool)

	for i := 0; i < MAX_RAW; i++ {
		for j := 0; j < MAX_COL; j++ {
			if board[j][i] == '.' {
				continue
			}
			_, exist := flags[board[j][i]]
			if exist {
				return false
			}
			flags[board[j][i]] = true
		}
		flags = make(map[byte]bool)
	}

	return true
}

func checkSquares(board [][]byte) bool {
	for i := 0; i < SQUARE_SIDE; i++ {
		for j := 0; j < SQUARE_SIDE; j++ {
			if !checkSqare(board, i, j) {
				return false
			}
		}
	}
	return true
}

func checkSqare(board [][]byte, rawSqare int, colSquare int) bool {
	START_RAW := rawSqare * SQUARE_SIDE
	START_COL := colSquare * SQUARE_SIDE

	var flags map[byte]bool = make(map[byte]bool)

	for i := START_RAW; i < START_RAW+SQUARE_SIDE; i++ {
		for j := START_COL; j < START_COL+SQUARE_SIDE; j++ {
			if board[i][j] == '.' {
				continue
			}
			_, exist := flags[board[i][j]]
			if exist {
				return false
			}
			flags[board[i][j]] = true
		}
	}

	return true
}
