package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	que := []*TreeNode{}
	que = append(que, root)
	for len(que) != 0 {
		root = que[0]
		if root.Left != nil {
			que = append(que, root.Left)
		}
		if root.Right != nil {
			que = append(que, root.Right)
		}
		f(que[0].Data)
		que = que[1:]

	}
}
