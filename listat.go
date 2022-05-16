package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	n := l
	cnt := 0
	if n == nil {
		return nil
	}
	if pos == 0 {
		return l
	}
	for n.Next != nil {
		cnt++
		if cnt == pos {
			return n.Next
		}
		n = n.Next

	}
	return nil
}
