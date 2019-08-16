package main

import (
	"fmt"
)

// Perm calls f with each permutation of a.
func PermInts(a []int, result [][]int, f func([]int)) {
	permInts(a, result, f, 0)
}

// Permute the values at index i to len(a)-1.
func permInts(a []int, result [][]int, f func([]int), i int) {
	cnt := 0

	if i > len(a) {
		f(a)
		cnt++
		result = append(result, a)
		//fmt.Println(a)
		return
	}
	permInts(a, result, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permInts(a, result, f, i+1)
		a[i], a[j] = a[j], a[i]
	}

}

func getIntFactorial(i int) int {
	ret := 1
	for cnt := i; cnt > 0; cnt-- {
		ret *= cnt
	}
	return ret
}

func main() {

	toPermutate := []int{1, 2, 3, 4}
	//permutations := [][]int{}
	permCount := getIntFactorial(len(toPermutate))
	permutations := make([][]int, permCount)

	for i := range permutations {
		permutations[i] = make([]int, len(toPermutate))
	}

	fmt.Println(&permutations)

	PermInts(toPermutate, permutations, func(a []int) {
		//fmt.Println(permutations)
		//intPermutations = append(intPermutations, a)
	})
	fmt.Println(permutations)
}
