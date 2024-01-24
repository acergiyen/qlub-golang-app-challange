package maxsumcalculator

import (
	"log"
)

// MaxSumCalculator represents a calculator type for calculating the maximum path sum.
type MaxSumCalculator struct {
	Logger *log.Logger // Logger object used for logging.
}

// NewMaxSumCalculator creates a new instance of MaxSumCalculator with a given log.Logger.
func NewMaxSumCalculator(logger *log.Logger) *MaxSumCalculator {
	return &MaxSumCalculator{Logger: logger}
}

// BinaryTree represents a node in a binary tree.
type BinaryTree struct {
	Value int         `json:"value"` // Represents the value of the node.
	Left  *BinaryTree `json:"left"`  // Represents the left subtree.
	Right *BinaryTree `json:"right"` // Represents the right subtree.
}

// TreeRequest represents the JSON structure containing information about the binary tree.
type TreeRequest struct {
	Tree struct {
		Nodes []struct {
			ID    string `json:"id"`
			Left  string `json:"left"`
			Right string `json:"right"`
			Value int    `json:"value"`
		} `json:"nodes"` // List containing tree nodes.
		Root string `json:"root"` // Identifier for the root node of the tree.
	} `json:"tree"`
}

// MaxPathSumResponse represents the JSON response containing the maximum path sum.
type MaxPathSumResponse struct {
	MaxPathSum int `json:"maxPathSum"`
}

// Handle calculates the maximum path sum for the given TreeRequest.
func (m *MaxSumCalculator) Handle(treeRequest TreeRequest) int {
	m.Logger.Println("Performing MaxPathSum calculation...")
	root := buildTree(treeRequest)
	result := maxPathSum(root)
	m.Logger.Printf("MaxPathSum result: %d\n", result)
	return result
}

// buildTree constructs a binary tree based on the information provided in the TreeRequest.
func buildTree(treeData TreeRequest) *BinaryTree {
	nodes := make(map[string]*BinaryTree)
	// Create binary tree nodes and store them in a map for easy lookup.
	for _, nodeData := range treeData.Tree.Nodes {
		node := &BinaryTree{Value: nodeData.Value}
		nodes[nodeData.ID] = node
	}

	// Connect the nodes to form the binary tree.
	for _, nodeData := range treeData.Tree.Nodes {
		node := nodes[nodeData.ID]
		if nodeData.Left != "" {
			node.Left = nodes[nodeData.Left]
		}
		if nodeData.Right != "" {
			node.Right = nodes[nodeData.Right]
		}
	}

	return nodes[treeData.Tree.Root] // Return the root of the binary tree.
}

// maxPathSum calculates the maximum path sum for the given binary tree.
func maxPathSum(root *BinaryTree) int {
	var maxSum int
	// Helper function is called to perform the recursive calculation.
	helper(root, &maxSum)
	return maxSum
}

// helper is a recursive function that helps in calculating the maximum path sum.
func helper(node *BinaryTree, maxSum *int) int {
	if node == nil {
		return 0 // Base case: an empty node contributes 0 to the path sum.
	}

	// Calculate the maximum path sum for the left and right subtrees.
	leftSum := max(helper(node.Left, maxSum), 0)
	rightSum := max(helper(node.Right, maxSum), 0)

	// Update the overall maximum path sum considering the current node.
	*maxSum = max(*maxSum, leftSum+rightSum+node.Value)

	// Return the maximum path sum considering the current node for further calculations.
	return max(leftSum, rightSum) + node.Value
}

// max is a utility function that returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
