package stack

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	var s Stack
	s.Values = []int{0, 1, 2, 3, 4, 5, 6, 7}
	got := s.Pop()
	want := 7
	if got != want {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
	t.Logf("excepted:%v, got:%v", want, got)
}

func TestStack_Peer(t *testing.T) {
	var s Stack
	s.Values = []int{0, 1, 2, 3, 4, 5, 6, 7}
	got := s.Peer()
	want := 7
	if got != want {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
	t.Logf("excepted:%v, got:%v", want, got)
}

func TestStack_Push(t *testing.T) {
	var s Stack
	s.Push(7)
	want := []int{7}
	if !reflect.DeepEqual(want, s.Values) {
		t.Errorf("excepted:%v, got:%v", want, s.Values)
	}
	t.Logf("excepted:%v, got:%v", want, s.Values)
}

func TestStack_Length(t *testing.T) {
	var s Stack
	s.Values = []int{0, 1, 2, 3, 4, 5, 6}
	got := s.Length()
	want := 7
	if got != want {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
	t.Logf("excepted:%v, got:%v", want, got)
}

func TestStack_IsEmpty(t *testing.T) {
	var s Stack
	got := s.IsEmpty()
	if !got {
		t.Errorf("excepted:%v, got:%v", false, got)
	}
	t.Logf("excepted:%v, got:%v", true, got)

}