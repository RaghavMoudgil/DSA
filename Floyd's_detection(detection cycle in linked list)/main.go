package main

import (
	"fmt"
)

// defininng a linked list in the struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func isCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func printList(head *ListNode) {
	first := head.Val
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next

		if isCycle(head) && first == head.Val {
			break
		}
	}

}

func main() {
	//creating a linked list here
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = head

	fmt.Print("this is the original list ")
	printList(head)
	fmt.Println("is it a cycle :", isCycle(head))
}
