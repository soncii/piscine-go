package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	var prev *NodeL
	for n != nil {
		if n.Data == data_ref {
			if n == l.Head {
				l.Head = n.Next
				n = n.Next
				continue
			}
			prev.Next = n.Next
			n = n.Next
			continue
		}
		prev = n
		n = n.Next

	}
}
