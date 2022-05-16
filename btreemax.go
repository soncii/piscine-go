package piscine

var (
	max   string
	mNode *TreeNode
)

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data > max {
		max = root.Data
		mNode = root
	}
	if root.Left != nil {
		BTreeMax(root.Left)
	}
	if root.Right != nil {
		BTreeMax(root.Right)
	}
	return mNode
}
