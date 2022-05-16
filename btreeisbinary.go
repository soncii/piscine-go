package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Data > root.Data {
			return false
		}
		return BTreeIsBinary(root.Left)
	}
	if root.Right != nil {
		if root.Right.Data <= root.Data {
			return false
		}
		return BTreeIsBinary(root.Right)
	}
	return true
}
