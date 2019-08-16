package main

import (
	"fmt"
)

func main() {
	generateBoards(4)
}

// multidimensional slice : https://go.googlesource.com/proposal/+/refs/heads/master/design/6282-table-data.md

// TODO : make this generate based off a board size
func generateRowPatterns(dimension int) [][]int {
	rp := make([][]int, 24)
	for i := 0; i < 24; i++ {
		rp[i] = make([]int, 4)
	}

	// TODO : Is there an easier way?
	copy(rp[0], []int{1, 2, 3, 4})
	copy(rp[1], []int{1, 2, 4, 3})
	copy(rp[2], []int{1, 3, 2, 4})
	copy(rp[3], []int{1, 3, 4, 2})
	copy(rp[4], []int{1, 4, 2, 3})
	copy(rp[5], []int{1, 4, 3, 2})
	copy(rp[6], []int{2, 1, 3, 4})
	copy(rp[7], []int{2, 1, 4, 3})
	copy(rp[8], []int{2, 3, 1, 4})
	copy(rp[9], []int{2, 3, 4, 1})
	copy(rp[10], []int{2, 4, 1, 3})
	copy(rp[11], []int{2, 4, 3, 1})
	copy(rp[12], []int{3, 1, 2, 4})
	copy(rp[13], []int{3, 1, 4, 2})
	copy(rp[14], []int{3, 2, 1, 4})
	copy(rp[15], []int{3, 2, 4, 1})
	copy(rp[16], []int{3, 4, 1, 2})
	copy(rp[17], []int{3, 4, 2, 1})
	copy(rp[18], []int{4, 1, 2, 3})
	copy(rp[19], []int{4, 1, 3, 2})
	copy(rp[20], []int{4, 2, 1, 3})
	copy(rp[21], []int{4, 2, 3, 1})
	copy(rp[22], []int{4, 3, 1, 2})
	copy(rp[23], []int{4, 3, 2, 1})

	return rp
}

type gameboard struct {
	dimension   int
	boardarray  [][]int
	hintPattern []int
}

// generateBoards will generate all valid boards of a given dimension
func generateBoards(dimension int) []gameboard {

	gameboards := []gameboard{} // now we have a slice to store the boards

	// TODO : find a way to generate this programatically.
	rowPatterns := generateRowPatterns(4)

	// Create a board slice
	var boardArray [][]int
	boardArray = make([][]int, dimension)

	for i := 0; i < dimension; i++ {
		boardArray[i] = make([]int, dimension)
	}
	// gameboard is now a slice with dimension x dimension size

	gameboardCount := 0
	skippedIterations := 0

	for r0 := 0; r0 < len(rowPatterns); r0++ {
		boardArray[0] = rowPatterns[r0]

		for r1 := 0; r1 < len(rowPatterns); r1++ {
			boardArray[1] = rowPatterns[r1]

			if checkBoardArray(&boardArray, dimension) == false {
				skippedIterations++
				boardArray[1] = []int{0, 0, 0, 0}
				continue
			}

			for r2 := 0; r2 < len(rowPatterns); r2++ {
				boardArray[2] = rowPatterns[r2]

				if checkBoardArray(&boardArray, dimension) == false {
					skippedIterations++
					boardArray[2] = []int{0, 0, 0, 0}
					continue
				}

				for r3 := 0; r3 < len(rowPatterns); r3++ {
					boardArray[3] = rowPatterns[r3]

					if checkBoardArray(&boardArray, dimension) == false {
						skippedIterations++
						boardArray[3] = []int{0, 0, 0, 0}
						continue
					}

					fmt.Printf("Complete Game Board!!! r1=%v r2=%v r3=%v\n", r1, r2, r3)
					gameboardCount++
					printBoard(&gameboard)

					gbs := gameboardstruct{dimension: 4, board: gameboard}
					gameboards := append(gameboards, gbs)

					generateHintPattern(&gameboard)
					fmt.Println()
				}
			}
		}
	}

	fmt.Printf("Gameboards Generated: %v\n", gameboardCount)
	fmt.Printf("Skipped Iterations: %v\n", skippedIterations)

	return gameboards
}

func clearBoard(gb *[4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			gb[i][j] = 0
		}
	}
}

func printBoard(gb *[4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", gb[i][j])
		}
		fmt.Println()
	}
}

// checkGameBoard will check a gameboard for validity, ignoring 0's
// it gets called as a game board is generated to determine if we are creating
// a bad board, that way we can short circuit the process and move to the next
// board

func checkBoardArray(gameboard *[][]int, dimension int) bool {

	// check each row for duplicate values
	for row := 0; row < dimension; row++ {
		for i := 0; i < (dimension - 1); i++ {
			val1 := gameboard[row][i]
			for j := i + 1; j < 4; j++ {
				val2 := gb[row][j]
				if val1 == val2 && val1 != 0 && val2 != 0 {
					return false
				}
			}
		}
	}

	// Check each column for any duplicate values
	for col := 0; col < 4; col++ {
		for i := 0; i < (4 - 1); i++ {
			val1 := gb[i][col]
			for j := i + 1; j < 4; j++ {
				val2 := gb[j][col]
				if val1 == val2 && val1 != 0 && val2 != 0 {
					return false
				}
			}
		}
	}

	return true
}

// doesHintPatternMatchBoard
func doesHintPatternMatchBoard(gb *[4][4]int, pattern [16]int) {

	p := generateHintPattern(gb)

}

// generateHintPattern
func generateHintPattern(gb *[4][4]int) [16]int {
	fmt.Println("Welcome to generateHint Pattern")
	fmt.Println("-------------------------------")
	fmt.Println("gameboard :")
	//printBoard(gb) //todo : remove after debugging

	ret := [16]int{}
	pos := 0

	// TODO : we could use a single loop to populate ret, refactor this later

	//fmt.Println("Top View :")
	// generate top
	for col := 0; col < 4; col++ {
		var buildingHeightArray [4]int
		for row := 0; row < 4; row++ {
			buildingHeightArray[row] = gb[row][col]
		}
		// Now we should have the buildingHeightArray built for this column looking from the top
		ret[pos] = visibleBuildingCount(buildingHeightArray)
		pos++
	}

	// generate right side
	for row := 0; row < 4; row++ {
		var buildingHeightArray [4]int
		cnt := 0
		for col := (4 - 1); col >= 0; col-- {
			buildingHeightArray[cnt] = gb[row][col]
			cnt++
		}
		ret[pos] = visibleBuildingCount(buildingHeightArray)
		pos++
	}

	// generate bottom
	for col := (4 - 1); col >= 0; col-- {
		var buildingHeightArray [4]int
		cnt := 0
		for row := (4 - 1); row >= 0; row-- {
			buildingHeightArray[cnt] = gb[row][col]
			cnt++
		}
		// Now we should have the buildingHeightArray built for this column looking from the top
		ret[pos] = visibleBuildingCount(buildingHeightArray)
		pos++
	}

	// generate left
	for row := (4 - 1); row >= 0; row-- {
		var buildingHeightArray [4]int
		cnt := 0
		for col := 0; col < 4; col++ {
			buildingHeightArray[cnt] = gb[row][col]
			cnt++
		}
		ret[pos] = visibleBuildingCount(buildingHeightArray)
		pos++
	}

	fmt.Printf("Hint Pattern : %v\n", ret)

	return ret
}

// visibleBuildingCount will count the number of visible buildings from the perspective
// ov the observer, the buildings parameter should be the building height in order from closest
// to furthest from the observer.
func visibleBuildingCount(buildingHeightArray [4]int) int {

	highestSeen := 0
	buildingsVisible := 0

	for x := 0; x < len(buildingHeightArray); x++ {
		if buildingHeightArray[x] > highestSeen {
			highestSeen = buildingHeightArray[x]
			buildingsVisible++
		}
	}

	return buildingsVisible
}

// TODO : generate a fn that will print the gameboard
// TODO : generate a hint pattern for each possible gameboard
// TODO : create a fn that tests a hint pattern and gets a group of possible boards
