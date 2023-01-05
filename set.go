package main

import "fmt"

var exists = struct{}{}

type setType = string

type set struct {
	m map[setType]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[setType]struct{})
	return s
}

func (s *set) Add(value setType) {
	s.m[value] = exists
}

func (s *set) Remove(value setType) {
	delete(s.m, value)
}

func (s *set) Contains(value setType) bool {
	_, c := s.m[value]
	return c
}
func (s *set) Len() int {
	return len(s.m)
}

func main() {
	s := NewSet()

	s.Add("Peter")
	s.Add("David")

	fmt.Println(s.Contains("Peter"))  // True
	fmt.Println(s.Contains("George")) // False

	s.Remove("David")
	fmt.Println(s.Contains("David")) // False
}
