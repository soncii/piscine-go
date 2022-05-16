package piscine

// type TreeNode struct {
// 	Data   string
// 	Left   *TreeNode
// 	Right  *TreeNode
// 	Parent *TreeNode
// }

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data == elem {
		return root
	}
	if root.Data > elem && root.Left != nil {
		return BTreeSearchItem(root.Left, elem)
	} else if root.Data < elem && root.Right != nil {
		return BTreeSearchItem(root.Right, elem)
	}
	return nil
}
