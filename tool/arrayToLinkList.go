package tool

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToLinkList(nums []int) *ListNode {

	listNode := &ListNode{}
	listHead := listNode
	for _, val := range nums {

		if listNode == nil {
			listNode = &ListNode{}
		}

		listNode.Val = val
		listNode = listNode.Next
	}

	return listHead
}

func LinkListToArray(linkHead *ListNode) []int {

	var arr []int

	for linkHead != nil {
		fmt.Println(linkHead.Val)
		arr = append(arr, linkHead.Val)
		linkHead = linkHead.Next
	}

	return arr
}
