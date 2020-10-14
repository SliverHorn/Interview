package main

import (
	"fmt"
)

type ListNode24 struct {
	Val  int
	Next *ListNode24
}

func main() {
	root := &ListNode24{
		Val: 1,
		Next: &ListNode24{
			Val: 2,
			Next: &ListNode24{
				Val: 3,
				Next: &ListNode24{
					Val: 4,
				},
			},
		},
	}
	a := swapPairs(root)
	fmt.Println(a)
}

func swapPairs(head *ListNode24) *ListNode24 {
	if head == nil || head.Next == nil {
		return head
	}
	firstNext := head.Next
	var first *ListNode24
	for head.Next != nil {
		headNext := head.Next
		if first != nil && first.Next != nil {
			first.Next = headNext
		}
		var next *ListNode24
		if head.Next != nil {
			next = head.Next.Next
		}
		if head.Next.Next != nil {
			head.Next = next
		} else {
			head.Next = nil
		}
		headNext.Next = head
		first = head
		if head.Next != nil {
			head = next
		}
	}
	return firstNext
}
