package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	var sb strings.Builder
	current := node
	for current != nil {
		if sb.Len() != 0 {
			sb.WriteString("->")
		}
		sb.WriteString(strconv.Itoa(current.Val))
		current = current.Next
	}
	return sb.String()
}

func NewList(values ...int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	current := head.Next
	for current != nil {
		if current.Val == prev.Val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return head
}
