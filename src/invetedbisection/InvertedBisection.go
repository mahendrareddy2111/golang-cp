package invetedbisection

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func InvertedBisection(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}
	mid, size := head.getMidAndSizeOfList()

	secondHalf := mid.Next

	if size%2 == 0 {
		head = head.Reverse(secondHalf)
		tempHead := secondHalf.Reverse(nil)
		temp := head

		for temp != nil && temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = tempHead
	} else {
		head = head.Reverse(mid)
		tempHead := secondHalf.Reverse(nil)
		temp := head

		for temp != nil && temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = mid
		mid.Next = tempHead
	}

	return head
}

func (l *LinkedList) Reverse(node *LinkedList) *LinkedList {
	//temp := l
	newHead := &LinkedList{l.Value, nil}
	l = l.Next
	for l != node {
		temp := l
		l = l.Next
		temp.Next = newHead
		newHead = temp
	}

	return newHead
}

func (l *LinkedList) getMidAndSizeOfList() (*LinkedList, int) {

	f, s := l, l
	size := 0

	for s != nil {
		s = s.Next
		size++
		if size%2 == 0 && s != nil {
			f = f.Next
		}

	}

	return f, size
}
