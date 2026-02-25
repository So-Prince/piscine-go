package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = listsort(n1)
	n2 = listsort(n2)

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	if n1.Data <= n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	} else {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}

func listsort(l *NodeI) *NodeI {
	head := l
	if head == nil {
		return nil
	}

	head.Next = listsort(head.Next)

	if head.Next != nil && head.Data > head.Next.Data {
		head = move(head)
	}

	return head
}
