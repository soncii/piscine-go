package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	d := n1
	for d.Next != nil {
		d = d.Next
	}
	d.Next = n2
	ListSort(n1)
	return n1
}
