package stack

type Stack struct {
	Values []int
}

func (s *Stack) Pop() int {
	if s.Length() == 0 {
		return 0
	}
	last := s.Values[s.Length()-1]
	s.Values = s.Values[:s.Length()-1]
	return last
}

func (s *Stack) Peer() int {
	return s.Values[s.Length()-1]
}

func (s *Stack) Push(v int) {
	s.Values = append(s.Values, v)
}

func (s *Stack) Length() int {
	return len(s.Values)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}
