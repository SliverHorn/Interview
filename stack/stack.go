package stack
/*

Stack:栈的顺序存储结构体
1. Pop: 从栈中弹出最后一个元素
2. Peer: 读取栈顶元素
3. Push: 向栈中添加元素
4. Length: 返回栈的长度
5. IsEmpty: 判断栈是否为空

*/

type StackInterface interface {
	Pop() int
	Peer() int
	Push(v int)
	Length() int
	IsEmpty() bool
}

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
