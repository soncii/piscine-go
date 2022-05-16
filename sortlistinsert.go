package piscine

// func listPushBack(l *NodeI, data int) *NodeI {
// 	n := &NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		l.Data = data_ref
		return l
	}
	var new NodeI
	new.Data = data_ref
	if l.Data > data_ref {
		new.Next = l
		return &new
	}
	head := l
	for l.Next != nil {
		if l.Next.Data > data_ref {
			new.Next = l.Next
			l.Next = &new
			return head
		}
		l = l.Next
	}
	l.Next = &new
	return head
}

// func SortListInsert(l *NodeI, data_ref int) *NodeI {
// 	listPushBack(l, data_ref)
// 	ListSort(l)
// 	return l
// }
