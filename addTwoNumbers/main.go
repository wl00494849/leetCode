package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	fmt.Scanln()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	newListNodeHead := newListNode

	for l1 != nil || l2 != nil {

		if newListNode != nil {
			if newListNode.Next == nil {
				newListNode.Next = &ListNode{}
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
			newListNode.Next.Next = &ListNode{
				Val: newListNode.Next.Val / 10,
			}

			newListNode.Next.Val = newListNode.Next.Val % 10
		}
		fmt.Println(newListNode.Val)
		newListNode = newListNode.Next
	}

	return newListNodeHead.Next
}
