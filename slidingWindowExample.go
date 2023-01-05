// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"fmt"
)

func main() {

	fmt.Println("findMaxConsecutiveOnes TC 1", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}) == 4)
	fmt.Println("findMaxConsecutiveOnes TC 1", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0}) == 4)
	fmt.Println("findMaxConsecutiveOnes TC 1", findMaxConsecutiveOnes([]int{1}) == 1)
	fmt.Println("findMaxConsecutiveOnes TC 1", findMaxConsecutiveOnes([]int{0}) == 1)
	fmt.Println("findMaxConsecutiveOnes TC 1", findMaxConsecutiveOnes([]int{0, 0}) == 1)
	fmt.Println("findMaxConsecutiveOnes TC 1", findMaxConsecutiveOnes([]int{1, 1}) == 2)
}

func findMaxConsecutiveOnes(nums []int) int {

	longestSequence := 0
	left, right := 0, 0
	num_zeroes := 0

	for right < len(nums) {

		if nums[right] == 0 {
			num_zeroes += 1
		}

		// if our window is invalid, contract our window
		for num_zeroes == 2 {
			if nums[left] == 0 {
				num_zeroes -= 1
			}
			left += 1
		}

		// update our longest sequence answer
		longestSequence = max(longestSequence, right-left+1)
		right += 1

	}
	return longestSequence

}
func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}
