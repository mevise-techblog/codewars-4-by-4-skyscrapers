package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

//GeneratePermutations will find all permutations for an array of integers and return the data on a channel
// one permutation at a time as they are discovered.
func GeneratePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}

func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

// TODO : make this generate based off a board size
func generateRowPatterns(dimension int) [][]int {

	cnt := 0

	permMePlease := []int{1, 2, 3, 4}
	permCount := factorial(len(permMePlease))

	permutations := make([][]int, permCount)
	for i := 0; i < 24; i++ {
		permutations[i] = make([]int, len(permMePlease))
	}

	for perm := range GeneratePermutations(permMePlease) {
		permutations[cnt] = perm
		//fmt.Println(perm)
		cnt++
	}

	return permutations
}

type gameboard struct {
	dimension   int
	boardarray  [][]int
	hintPattern []int
}

// generateBoards will generate all valid boards of a given dimension
func generateBoards(dimension int) []*gameboard {

	gameboards := []*gameboard{} // now we have a slice to store the boards

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

			if checkBoardArray(boardArray, dimension) == false {
				skippedIterations++
				boardArray[1] = []int{0, 0, 0, 0}
				continue
			}

			for r2 := 0; r2 < len(rowPatterns); r2++ {
				boardArray[2] = rowPatterns[r2]

				if checkBoardArray(boardArray, dimension) == false {
					skippedIterations++
					boardArray[2] = []int{0, 0, 0, 0}
					continue
				}

				for r3 := 0; r3 < len(rowPatterns); r3++ {
					boardArray[3] = rowPatterns[r3]

					if checkBoardArray(boardArray, dimension) == false {
						skippedIterations++
						boardArray[3] = []int{0, 0, 0, 0}
						continue
					}

					fmt.Printf("Complete Game Board!!! r1=%v r2=%v r3=%v\n", r1, r2, r3)
					gameboardCount++
					printBoard(boardArray)

					gbs := gameboard{dimension: dimension, boardarray: boardArray}
					gameboards := append(gameboards, &gbs)

					//generateHintPattern(gameboard)
					fmt.Println(gameboards)
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

func printBoard(gb [][]int) {
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

func checkBoardArray(ba [][]int, dimension int) bool {

	// check each row for duplicate values
	for row := 0; row < dimension; row++ {
		for i := 0; i < (dimension - 1); i++ {
			val1 := ba[row][i]
			for j := i + 1; j < 4; j++ {
				val2 := ba[row][j]
				if val1 == val2 && val1 != 0 && val2 != 0 {
					return false
				}
			}
		}
	}

	// Check each column for any duplicate values
	for col := 0; col < 4; col++ {
		for i := 0; i < (4 - 1); i++ {
			val1 := ba[i][col]
			for j := i + 1; j < 4; j++ {
				val2 := ba[j][col]
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
