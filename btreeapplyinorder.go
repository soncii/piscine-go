package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                stri
// }

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	if root != nil {
		f(root.Data)
	}
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}
