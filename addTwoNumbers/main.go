package main

import (
	"exam/tool"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	num1 := []int{2, 3, 4}
	num2 := []int{5, 6, 4}

	l1 := tool.ArrayToLinkList(num1)
	l2 := tool.ArrayToLinkList(num2)

	listHead := addTwoNumbers(l1, l2)

	res := tool.LinkListToArray(listHead)

	fmt.Printf("result : %+v", res)
}

func addTwoNumbers(l1 *tool.ListNode, l2 *tool.ListNode) *tool.ListNode {
	newListNode := &tool.ListNode{}
	newListNodeHead := newListNode

	for l1 != nil || l2 != nil {

		if newListNode != nil {
			if newListNode.Next == nil {
				newListNode.Next = &tool.ListNode{}
			}
		}

		if l1 != nil {
			newListNode.Next.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			newListNode.Next.Val += l2.Val
			l2 = l2.Next
		}

		if newListNode.Next.Val > 9 {
			newListNode.Next.Next = &tool.ListNode{
				Val: newListNode.Next.Val / 10,
			}

			newListNode.Next.Val = newListNode.Next.Val % 10
		}
		newListNode = newListNode.Next
	}

	return newListNodeHead.Next
}
