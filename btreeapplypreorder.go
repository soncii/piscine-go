package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                stri
// }

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
	}
	if root.Left != nil {
		BTreeApplyPreorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyPreorder(root.Right, f)
	}
}
