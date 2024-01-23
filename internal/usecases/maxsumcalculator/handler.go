package maxsumcalculator

import (
	"log"
)

type MaxSumCalculator struct {
	Logger *log.Logger
}

func NewMaxSumCalculator(logger *log.Logger) *MaxSumCalculator {
	return &MaxSumCalculator{Logger: logger}
}

type BinaryTree struct {
	Value int         `json:"value"`
	Left  *BinaryTree `json:"left"`
	Right *BinaryTree `json:"right"`
}

type TreeRequest struct {
	Tree struct {
		Nodes []struct {
			ID    string `json:"id"`
			Left  string `json:"left"`
			Right string `json:"right"`
			Value int    `json:"value"`
		} `json:"nodes"`
		Root string `json:"root"`
	} `json:"tree"`
}

type MaxPathSumResponse struct {
	MaxPathSum int `json:"maxPathSum"`
}

func (m *MaxSumCalculator) Handle(treeRequest TreeRequest) int {
	m.Logger.Println("Handling MaxPathSum calculation...")
	root := buildTree(treeRequest)
	result := maxPathSum(root)
	m.Logger.Printf("MaxPathSum result: %d\n", result)
	return result
}
func buildTree(treeData TreeRequest) *BinaryTree {
	nodes := make(map[string]*BinaryTree)
	for _, nodeData := range treeData.Tree.Nodes {
		node := &BinaryTree{Value: nodeData.Value}
		nodes[nodeData.ID] = node
	}

	for _, nodeData := range treeData.Tree.Nodes {
		node := nodes[nodeData.ID]
		if nodeData.Left != "" {
			node.Left = nodes[nodeData.Left]
		}
		if nodeData.Right != "" {
			node.Right = nodes[nodeData.Right]
		}
	}

	return nodes[treeData.Tree.Root]
}

func maxPathSum(root *BinaryTree) int {
	var maxSum int
	helper(root, &maxSum)
	return maxSum
}

func helper(node *BinaryTree, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftSum := max(helper(node.Left, maxSum), 0)
	rightSum := max(helper(node.Right, maxSum), 0)

	*maxSum = max(*maxSum, leftSum+rightSum+node.Value)

	return max(leftSum, rightSum) + node.Value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
