package set

// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	columns := [9][9]bool{}
	partialSudoku := [9][9]bool{}

	rowBucket := 0
	for i := 0; i < 9; i++ {
		columnBucket := 0
		for j := 0; j < 9; j++ {
			if i != 0 && j == 0 && i%3 == 0 {
				rowBucket += 3
			}
			if j != 0 && j%3 == 0 {
				columnBucket++
			}
			if board[i][j] == '.' {
				continue
			}

			if rows[i][int(board[i][j]-'0')-1] || columns[j][int(board[i][j]-'0')-1] ||
				partialSudoku[rowBucket+columnBucket][int(board[i][j]-'0')-1] {
				return false
			} else {
				rows[i][int(board[i][j]-'0')-1] = true
				columns[j][int(board[i][j]-'0')-1] = true
				partialSudoku[rowBucket+columnBucket][int(board[i][j]-'0')-1] = true
			}
		}
	}

	return true
}
