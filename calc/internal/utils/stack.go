package utils

type Stack []interface{}

func CreateStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Stack) Top() interface{} {
	if !s.IsEmpty() {
		return (*s)[len(*s)-1]
	}
	return nil
}
