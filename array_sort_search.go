package main

import "sort"

func sortAndSearch() {

	arr := []int{2, 3, 5, 4, 2, 1}

	sort.Ints(arr)

	toSearch(&arr, 3)

}

func toSearch(arr *[]int, search int) int {

	low := 0
	high := len(*arr) - 1

	mid := high / 2

	for high-low > 1 {

		if (*arr)[mid] == search {

			return mid

		} else if (*arr)[mid] < search {

			low = mid
		} else {

			high = mid
		}
		mid = (high + low) / 2

	}

	for i := low; i <= high; i++ {
		if (*arr)[i] == search {
			return i
		}
	}

	return -1

}
