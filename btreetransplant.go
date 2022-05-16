package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		rplc.Parent = root.Parent
		if root.Parent.Left == node {
			root.Parent.Left = rplc
		}
		if root.Parent.Right == node {
			root.Parent.Right = rplc
		}
		root = rplc
		return root
	}
	if root.Left != nil {
		BTreeTransplant(root.Left, node, rplc)
	}
	if root.Right != nil {
		BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
