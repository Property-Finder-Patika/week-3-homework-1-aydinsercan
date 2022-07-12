package main

import (
	"fmt"
	"math/rand"
)

//Given an array, rotate the array to the right or to the left by k steps, where k is non-negative.
//O(n) solution

func main() {

	argArr := rand.Perm(5)
	fmt.Println("Original Array : ", argArr)

	fmt.Println("Right Rotate Result : ", RightRotateSlice(argArr, 3))
	fmt.Println("Left Rotate Result : ", LeftRotateSlice(argArr, 3))

}

func RightRotateSlice(slc []int, step int) []int {

	var last int
	for i := 0; i < step; i++ {
		last, slc = slc[len(slc)-1], slc[:len(slc)-1]
		//Prepend an integer to this slice
		slc = append([]int{last}, slc...)
		fmt.Printf("Step %d \n", i+1)
		fmt.Println(slc)
	}

	return slc
}

func LeftRotateSlice(slc []int, step int) []int {

	var front int
	for i := 0; i < step; i++ {
		front, slc = slc[0], slc[1:]
		//append an integer to this slice
		slc = append(slc, front)
		fmt.Printf("Step %d \n", i+1)
		fmt.Println(slc)
	}
	return slc

}

/* Example Output

Orjinal Array :  [0 4 2 3 1]
Step 1
[1 0 4 2 3]
Step 2
[3 1 0 4 2]
Step 3
[2 3 1 0 4]
Right Rotate Result :  [2 3 1 0 4]
Step 1
[4 2 3 1 0]
Step 2
[2 3 1 0 4]
Step 3
[3 1 0 4 2]
Left Rotate Result :  [3 1 0 4 2]
*/
