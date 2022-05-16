package piscine

func ListReverse(l *List) {
	var prev *NodeL
	cur := l.Head
	var nex *NodeL
	for cur != nil {
		nex = cur.Next
		cur.Next = prev
		prev = cur
		cur = nex
	}
	l.Head = l.Tail
}

// func ListReverse(l *List) {
// 	n := ListSize(l)
// 	arr := make([]interface{}, n)
// 	it := l.Head
// 	for i := 0; i < n; i++ {
// 		arr[i] = it.Data
// 		it = it.Next
// 		//	fmt.Print(arr[i])
// 	}
// 	ListClear(l)
// 	for i := 0; i < n; i++ {
// 		ListPushFront(l, arr[i])
// 	}
// }

// func ListReverse(l *List) {
// 	var d *List
// 	n := l.Head
// 	for n.Next != nil {
// 		ListPushFront(d, n.Data)
// 		n = n.Next
// 	}
// 	d.Tail = nil
// }
