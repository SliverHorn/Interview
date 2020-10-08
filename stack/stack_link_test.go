package stack

import "testing"

func TestStackLink_Pop(t *testing.T) {
	var sl StackLink
	sl.Push(0)
	sl.Push(1)
	sl.Push(2)
	sl.Push(3)
	sl.Push(4)
	sl.Push(5)
	sl.Push(6)
	sl.Push(7)
	t.Logf("sl len:%v\n", sl.Len)
	got := sl.Pop()
	want := 7
	if got != want && sl.Len != 7 {
		t.Errorf("want:%v, got:%v, len:%v\n",want, got, sl.Len)
	}
	t.Logf("want:%v, got:%v, len:%v\n",want, got, sl.Len)
}

func TestStackLink_Peer(t *testing.T) {
	var sl StackLink
	sl.Push(0)
	sl.Push(1)
	sl.Push(2)
	sl.Push(3)
	sl.Push(4)
	sl.Push(5)
	sl.Push(6)
	sl.Push(7)
	t.Logf("sl len:%v\n", sl.Len)
	got := sl.Peer()
	want := 7
	if got != want {
		t.Errorf("want:%v, got:%v, len:%v\n",want, got, sl.Len)
	}
	t.Logf("want:%v, got:%v, len:%v\n",want, got, sl.Len)
}

func TestStackLink_Push(t *testing.T) {
	var sl StackLink
	sl.Push(7)
	if sl.Len == 0 && sl.Top == nil {
		t.Errorf("want:len=%v, top=%v, got:len=%v, top=%v",0, nil, sl.Len, sl.Top)
	}
	t.Logf("want:len=%v, top=%v, got:len=%v, top=%v",1, "&{7 <nil>}", sl.Len, sl.Top)
}

func TestStackLink_Length(t *testing.T) {
	var sl StackLink
	sl.Push(1)
	sl.Push(2)
	sl.Push(3)
	sl.Push(4)
	sl.Push(5)
	sl.Push(6)
	sl.Push(7)
	t.Logf("sl len:%v\n", sl.Len)
	if sl.Len != 7 {
		t.Errorf("want:%v, got:%v\n",7, sl.Len)
	}
	t.Logf("want:%v, got:%v\n",7, sl.Len)

}

func TestStackLink_IsEmpty(t *testing.T) {
	var sl StackLink
	sl.Push(7)
	got := sl.IsEmpty()
	want := false
	if got != want {
		t.Errorf("want:%v, got:%v\n",want, got)
	}
	t.Logf("want:%v, got:%v\n",want, got)

}