package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node7 := TreeNode{
		Val: 7,
	}
	node2 := TreeNode{
		Val: 2,
	}
	node11 := TreeNode{
		Val:   11,
		Left:  &node7,
		Right: &node2,
	}
	nodeLeft4 := TreeNode{
		Val:  4,
		Left: &node11,
	}

	node1 := TreeNode{
		Val: 1,
	}
	nodeRight4 := TreeNode{
		Val:   4,
		Right: &node1,
	}
	node13 := TreeNode{
		Val: 13,
	}
	node8 := TreeNode{
		Val:   8,
		Left:  &node13,
		Right: &nodeRight4,
	}
	root := TreeNode{
		Val:   5,
		Left:  &nodeLeft4,
		Right: &node8,
	}

	ans := goodNodes(&root)
	fmt.Print(ans)
}

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt)
}

func dfs(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	ans := 0
	fnMax := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	left := dfs(node.Left, fnMax(node.Val, max))
	right := dfs(node.Right, fnMax(node.Val, max))

	ans = left + right
	if node.Val >= max {
		ans += 1
	}

	return ans
}
