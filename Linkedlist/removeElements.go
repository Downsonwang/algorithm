package linkedlist

// leetcode 203. Remove LinkedList Elements

type ListAllNode struct {
	Val  int
	Next *ListAllNode
}

func removeElements(head *ListAllNode, val int) *ListAllNode {

	if head == nil {
		return &ListAllNode{}
	}

	prev := &ListAllNode{}
	prev.Next = head

	cur := prev

	for cur != nil && cur.Next != nil {
		//确保本节点和下一个结点都不为空
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return prev.Next
}
