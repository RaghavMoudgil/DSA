package main

import (
	"fmt"
)

// defininng a linked list in the struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func revrseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		next := head.Next
		head.Next = rev
		rev = head
		head = next
	}
	return rev
}

func printList(head *ListNode) {

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	//creating a linked list here
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Print("this is the original list ")
	printList(head)
	fmt.Print("below is the reversed list ")
	printList(revrseList(head))
}
