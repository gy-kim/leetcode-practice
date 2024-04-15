package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	m := map[int]*ListNode{}

	for i := 4; i >= 0; i-- {
		next := m[i+1]
		node := &ListNode{
			Val: i,
		}
		if next != nil {
			node.Next = next
		}
		m[i] = node
	}

	// for k, v := range m {
	// 	fmt.Printf("%i - %#v\n", k, v)
	// }

	getMiddle(m[0])
}

func getMiddle(head *ListNode) {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Println(slow.Val)
}
