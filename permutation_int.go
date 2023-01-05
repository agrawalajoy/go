	package main

import "fmt"

func main() {

	permutation(4)
}

func permutation(n int) {

	// for i := 0; i < n; i++ {
	arr := make([]bool, n)
	outArr := make([]int, n)
	// arr[i] = false
	// outArr[0] = i
	getNextPermutation(&arr, &outArr, 0, n)

	// }

}

func getNextPermutation(boolArr *[]bool, outArr *[]int, currentIndex int, totalIndex int) {

	if currentIndex == totalIndex {
		fmt.Println("F0und One ", outArr)
		return
	}
	for i := 0; i < totalIndex; i++ {

		if !(*boolArr)[i] {
			(*boolArr)[i] = true
			(*outArr)[currentIndex] = i
			getNextPermutation(boolArr, outArr, currentIndex+1, totalIndex)
			(*boolArr)[i] = false
			(*outArr)[currentIndex] = 0
		}

	}
}

//  1 2 3 4 5
//  1 3 2 4 5
//  1 4
