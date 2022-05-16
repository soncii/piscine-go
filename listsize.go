package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	n := l.Head
	cnt := 0
	for n.Next != nil {
		cnt++
		n = n.Next
	}
	return cnt + 1
}
