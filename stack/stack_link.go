package stack

/*

StackLink:栈的链式存储结构体
1. Pop: 从栈中弹出最后一个元素
2. Peer: 读取栈顶元素
3. Push: 向栈中添加元素
4. Length: 返回栈的长度
5. IsEmpty: 判断栈是否为空

 */

type StackLinkInterface interface {
	Pop() interface{}
	Peer() interface{}
	Push(v interface{})
	Length() int
	IsEmpty() bool
}

type StackLink struct {
    Top *Node
    Len int
}

type Node struct {
    Value interface{}
	Pointer *Node
}

func (l *StackLink) Pop() interface{} {
	topValue := l.Top.Value
	l.Top = l.Top.Pointer
	l.Len--
	return topValue
}

func (l *StackLink) Peer() interface{} {
	return l.Top.Value
}

func (l *StackLink) Push(v interface{}) {
	node := &Node{Value: v, Pointer: l.Top}
	l.Top = node
	l.Len++
}

func (l *StackLink) Length() int {
	return l.Len
}

func (l *StackLink) IsEmpty() bool {
	return l.Len == 0
}
