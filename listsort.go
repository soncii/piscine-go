// MERGESORT O(n^2)
// package piscine
// type NodeI struct {
// 	Data int
// 	Next *NodeI
// }

// var sorted *NodeI = nil

// func sortedInsert(new *NodeI) {
// 	if sorted == nil || sorted.Data >= new.Data {
// 		new.Next = sorted
// 		sorted = new
// 	} else {
// 		current := sorted
// 		for current.Next != nil && current.Next.Data < new.Data {
// 			current = current.Next
// 		}
// 		new.Next = current.Next
// 		current.Next = new
// 	}
// }

// func ListSort(l *NodeI) *NodeI {
// 	current := l
// 	for current != nil {
// 		next := current.Next
// 		sortedInsert(current)
// 		current = next
// 	}

// 	return sorted
// }

package piscine

import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	start := true
	var head *NodeI

	if l == nil || l.Next == nil {
		return l
	}

	for l != nil {
		next := l.Next
		if start {
			start = false
			head = l
		}

		for next != nil {
			if l.Data > next.Data {
				l.Data, next.Data = next.Data, l.Data
			}
			next = next.Next
		}
		l = l.Next
	}
	return head
}

func ListSort2(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	n := l
	a := []int{}
	for l != nil {
		a = append(a, l.Data)
		l = l.Next
	}

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	head := n
	for i := 0; i < len(a); i++ {
		n.Data = a[i]
		if i == 0 {
			head = n
		}
		fmt.Println(n)
		n = n.Next
	}
	return head
}
