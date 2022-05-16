package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	R := BTreeLevelCount(root.Right)
	L := BTreeLevelCount(root.Left)

	if L > R {
		return L + 1
	}
	return R + 1
}
