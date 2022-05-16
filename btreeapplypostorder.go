package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                stri
// }

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyPostorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyPostorder(root.Right, f)
	}
	if root != nil {
		f(root.Data)
	}
}
