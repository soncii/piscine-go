package piscine

var min = "i"

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data < min {
		min = root.Data
		mNode = root
	}
	if root.Left != nil {
		BTreeMin(root.Left)
	}
	if root.Right != nil {
		BTreeMin(root.Right)
	}
	return mNode
}
