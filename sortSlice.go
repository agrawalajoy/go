package main

import (
	"fmt"
	"sort"
)

//Sort on struct
type employee struct {
	name   string
	age    int
	number int
}

func main() {

	employees := []employee{
		{"ajoy", 1, 3},
		{"ajoy65", 6, 5},
		{"ajoy55", 5, 5},
		{"ajoy25", 2, 6},
		{"ajoy75", 7, 5},
		{"ajoy35", 8, 7},
		{"ajoy85", 8, 8},
		{"ajoy45", 4, 5},
	}
	fmt.Println("Before Sort", employees)

	sort.Slice(employees, func(i, j int) bool {
		// return (employees[i].age + employees[i].number) > (employees[j].age + employees[j].number)

		if employees[i].age == employees[j].age {
			return employees[i].number < employees[j].number
		}

		return employees[i].age < employees[j].age
	})

	fmt.Println("After Sort", employees)

	/// SEARCH
	age := 6
	index := sort.Search(len(employees), func(index int) bool { return employees[index].age >= age })
	if index >= len(employees) || employees[index].age != age {
		fmt.Println("NOT FOUND index=>", index)
	} else {
		fmt.Println("  FOUND index=>", index, employees[index])
	}

}
