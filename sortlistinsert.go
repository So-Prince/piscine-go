package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = listPushBackNode(l, data_ref)
	l = listSort(l)
	return l
}

func listPushBackNode(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data, Next: nil}

	if l == nil {
		return n
	}

	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = n
	return l
}

func listSort(l *NodeI) *NodeI {
	head := l
	if head == nil {
		return nil
	}

	head.Next = listSort(head.Next)

	if head.Next != nil && head.Data > head.Next.Data {
		head = move(head)
	}

	return head
}
