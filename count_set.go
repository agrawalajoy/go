package main

import "fmt"

type setType = string

type countSet struct {
	m map[setType]int
}

func NewSet() *countSet {
	s := &countSet{}
	s.m = make(map[setType]int)
	return s
}

func (s *countSet) Add(value setType) {
	if v, ok := s.m[value]; ok {
		s.m[value] = v + 1
	} else {
		s.m[value] = 1
	}
}

func (s *countSet) Remove(value setType) {

	if v, ok := s.m[value]; ok {
		if v <= 1 {
			delete(s.m, value)
		} else {
			s.m[value] = v - 1
		}
	}
}

func (s *countSet) Contains(value setType) bool {
	_, c := s.m[value]
	return c
}

func (s *countSet) GetCount(value setType) int {
	if v, ok := s.m[value]; ok {
		return v
	}
	return 0
}

func (s *countSet) Len(value setType) int {
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
