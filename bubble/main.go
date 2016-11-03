package main

import (
	"fmt"
)

// I mean, this is a sort, if there are only two items, right? Order here (and everywhere else) means that the earlier items are smaller (ascending?)
func SwapFirstTwoElementsIfOutOfOrder(slice []int) []int {
	return slice
}

// Check the test cases if this doesnt make sense.
//
// [x, a, b, c, d] -> [a, b, c, x, d] (if d is the only item that is bigger than X)
func MoveFirstElementUntilAllItemsBeforeItAreSmaller(slice []int) []int {
	return slice
}

// Sorted as in [0, 1, 2, 3]
func KeepMovingTheFirstElementUntilTheSliceIsSorted(slice []int) []int {
	return slice
}

func BubbleSort(slice []int) []int {
	return slice
}

func main() {
	fmt.Println("Test stuff:", SwapFirstTwoElementsIfOutOfOrder([]int{0, 1}))
}
