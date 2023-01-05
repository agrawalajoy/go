package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{-1, 2, 1, -4}
	nums := []int{0, 1, 2}
	target := 3

	fmt.Println("closet to target is threeSumClosest", threeSumClosest(nums, target))

}
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	// fmt.Println("Nums", nums)

	tripletSum := nums[0] + nums[1] + nums[2]

	diff := abs(target - tripletSum)
	lo := 0
	hi := 0
	lTarget := 0

	for i := 0; i < len(nums)-2; i++ {

		lo = i + 1
		hi = len(nums) - 1

		lTarget = target - nums[i]

		for lo < hi {

			if nums[lo]+nums[hi] == lTarget {
				return target
			}

			if abs(lTarget-nums[lo]-nums[hi]) < diff {
				diff = abs(lTarget - nums[lo] - nums[hi])
				tripletSum = nums[lo] + nums[hi] + nums[i]
			}

			if nums[lo]+nums[hi] < lTarget {
				lo++
			} else {
				hi--
			}
		}

	}

	return tripletSum
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
