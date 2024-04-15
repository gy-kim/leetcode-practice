package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	targetSum := 22

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
	node4 := TreeNode{
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
	node5 := TreeNode{
		Val:   5,
		Left:  &node4,
		Right: &node8,
	}

	ans := hasPathSum(&node5, targetSum)
	fmt.Println(ans)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, curr, target int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return (curr + node.Val) == target
	}

	curr += node.Val

	left := dfs(node.Left, curr, target)
	right := dfs(node.Right, curr, target)

	return left || right
}
