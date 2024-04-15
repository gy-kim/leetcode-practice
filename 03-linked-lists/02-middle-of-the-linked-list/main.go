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

	ans := middleNode(m[1])
	fmt.Println(ans)

}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
