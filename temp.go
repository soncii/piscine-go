package piscine

// type TreeNode struct {
// 	Data   string
// 	Left   *TreeNode
// 	Right  *TreeNode
// 	Parent *TreeNode
// }

// var (
// 	max   string
// 	mNode *TreeNode
// )

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if root.Data < data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	return root
// }
// var (
// 	min   = "i"
// 	mNode *TreeNode
// )

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if root.Data < data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	return root
// }

// var L, R int = 0, 0

// func BTreeLevelCount(root *TreeNode) int {
// 	if root == nil {
// 		return 1
// 	}
// 	if root.Left != nil {
// 		L = L + 1
// 		BTreeLevelCount(root.Left)
// 	}
// 	if root.Right != nil {
// 		R = R + 1
// 		BTreeLevelCount(root.Right)
// 	}
// 	if L > R {
// 		return L + 1
// 	}
// 	return R + 1
// }
// func BTreeIsBinary(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if root.Left != nil {
// 		if root.Left.Data < root.Data {
// 			return false
// 		}
// 		return BTreeIsBinary(root.Left)
// 	}
// 	if root.Right != nil {
// 		if root.Right.Data > root.Data {
// 			return false
// 		}
// 		return BTreeIsBinary(root.Right)
// 	}
// 	return true
// }

// func BTreeMax(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if root.Data > max {
// 		max = root.Data
// 		mNode = root
// 	}
// 	if root.Left != nil {
// 		BTreeMax(root.Left)
// 	}
// 	if root.Right != nil {
// 		BTreeMax(root.Right)
// 	}
// 	return mNode
// }

// func BTreeMin(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if root.Data < min {
// 		min = root.Data
// 		mNode = root
// 	}
// 	if root.Left != nil {
// 		BTreeMin(root.Left)
// 	}
// 	if root.Right != nil {
// 		BTreeMin(root.Right)
// 	}
// 	return mNode
// }

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeInsertData(root, "6")
// 	BTreeInsertData(root, "9")
// 	//	max := BTreeMin(root)
// 	//	fmt.Println(max.Data)
// }
