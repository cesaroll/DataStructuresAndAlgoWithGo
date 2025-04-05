package main

import "fmt"

func main() {
	head := createLinkedList(1, 2, 3, 4, 5)

	printLinkedList(head)
	head = reverseLinkedList(head)
	printLinkedList(head)
}

type Node struct {
	Value int
	Next  *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func createLinkedList(nums ...int) *Node {
	head := &Node{Value: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &Node{Value: nums[i]}
		curr = curr.Next
	}
	return head
}

func printLinkedList(head *Node) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}
}
