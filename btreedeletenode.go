package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if root.Left == nil && root.Right == nil {
			if root.Parent.Left == node {
				root.Parent.Left = nil
			}
			if root.Parent.Right == node {
				root.Parent.Right = nil
			}

		} else if root.Left == nil && root.Right != nil {
			if root.Parent.Left == node {
				root.Parent.Left = root.Right
			}
			if root.Parent.Right == node {
				root.Parent.Right = root.Right
			}

		} else if root.Left != nil && root.Right == nil {
			if root.Parent.Left == node {
				//	root.Left.Parent = root.Parent.Left
				root.Parent.Left = root.Left
			}
			if root.Parent.Right == node {
				//	root.Right.Parent = root.Parent.Right
				root.Parent.Right = root.Left
			}
		} else {
			min := BTreeMin(root.Right)
			BTreeDeleteNode(root, min)
			min.Right = root.Right
			min.Left = root.Left
			if root.Parent != nil {
				min.Parent = root.Parent
			}
			root = min

			// 	if min != node.Right {
			// 		BTreeTransplant(root, min, min.Right)
			// 		min.Right = root.Right
			// 		min.Right.Parent = min
			// 	}
			// 	BTreeTransplant(root, node, min)
			// 	min.Left = node.Left
			// 	min.Left.Parent = min
			// }
		}
	}
	if root.Left != nil {
		BTreeDeleteNode(root.Left, node)
	}
	if root.Right != nil {
		BTreeDeleteNode(root.Right, node)
	}
	return root
}

// func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if node.Left == nil {
// 		BTreeTransplant(root, node, node.Right)
// 	} else if node.Right == nil {
// 		BTreeTransplant(root, node, node.Left)
// 	} else {
// 		min := BTreeMin(root.Right)
// 		BTreeDeleteNode(root, min)
// 		if min != node.Right {
// 			BTreeTransplant(root, min, min.Right)
// 			min.Right = node.Right
// 			min.Right.Parent = min
// 		}
// 		BTreeTransplant(root, node, min)
// 		min.Left = node.Left
// 		min.Left.Parent = min
// 	}
// 	return root
// }
