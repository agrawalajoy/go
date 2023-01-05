type Stack []YourType

func (s *Stack) Push(v YourType) {
	*s = append(*s, v)
}

func (s *Stack) Pop() YourType {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}
