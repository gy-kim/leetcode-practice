package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	m := map[int]*ListNode{
		3: &ListNode{
			Val: 3,
		},
		2: &ListNode{
			Val: 2,
		},
		0: &ListNode{
			Val: 0,
		},
		-4: &ListNode{
			Val: -4,
		},
	}

	m[3].Next = m[2]
	m[2].Next = m[0]
	m[0].Next = m[-4]
	m[-4].Next = m[2]

	// for k, v := range m {
	// 	fmt.Printf("%i - %#v\n", k, v)
	// }

	ans := hasCycle(m[3])

	// ans := detectCycle(m[3])
	fmt.Println(ans)

}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// func detectCycle(head *ListNode) bool {
// 	m := map[*ListNode]struct{}{}

// 	for head != nil {
// 		if _, ok := m[head]; ok {
// 			return true
// 		}

// 		m[head] = struct{}{}
// 		head = head.Next
// 	}

// 	return false
// }
