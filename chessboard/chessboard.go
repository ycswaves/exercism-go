package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0
	for _, occupied := range cb[rank] {
		if occupied {
			count++
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0
	for _, ranks := range cb {
		for fileIndex, occupied := range ranks {
			if fileIndex == file-1 && occupied {
				count++
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, ranks := range cb {
		for range ranks {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, ranks := range cb {
		for _, occupied := range ranks {
			if occupied {
				count++
			}
		}
	}

	return count
}
