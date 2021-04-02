package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var node TreeNode
	node.Val = -1
	node.Left = root
	dfs(&node)
	return node.Left
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left)
	r := dfs(root.Right)
	if l == 0 {
		root.Left = nil
	}

	if r == 0 {
		root.Right = nil
	}

	return root.Val + l + r
}
