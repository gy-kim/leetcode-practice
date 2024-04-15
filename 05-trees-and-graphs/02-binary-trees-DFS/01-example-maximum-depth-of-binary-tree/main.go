package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node7 := TreeNode{
		Val: 7,
	}
	node15 := TreeNode{
		Val: 15,
	}

	node20 := TreeNode{
		Val:   20,
		Left:  &node15,
		Right: &node7,
	}

	node9 := TreeNode{
		Val: 9,
	}

	node3 := TreeNode{
		Val:   3,
		Left:  &node9,
		Right: &node20,
	}

	ans := maxDepth(&node3)
	fmt.Println(ans)

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	ans := max(left, right) + 1
	return ans
}
