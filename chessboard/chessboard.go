package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	for K, v := range cb {
		if K == file {
			for _, v := range v {
				if v {
					count++
				}
			}
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0

	for _, v := range cb {
		for k, v := range v {
			if k+1 == rank {
				if v {
					count++
				}
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	file := 0
	rank := 0
	count := 0

	for k := range cb {
		for range k {
			rank++
		}

		file++
	}

	count = file * rank

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for k := range cb {
		count += CountInFile(cb, k)
	}

	return count
}
