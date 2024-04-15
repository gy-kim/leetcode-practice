package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	m := map[int]*ListNode{}

	for i := 5; i >= 1; i-- {
		next := m[i+1]
		node := &ListNode{
			Val: i,
		}
		if next != nil {
			node.Next = next
		}
		m[i] = node
	}

	findNode(m[1], 2)

}

func findNode(head *ListNode, k int) {
	slow := head
	fast := head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fmt.Println(slow.Val)
}
