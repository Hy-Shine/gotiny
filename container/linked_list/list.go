package linkedlist

type Node[K any] struct {
	Data K
	Next *Node[K]
}

func New[K any]() *Node[K] {
	return &Node[K]{}
}

func (ll *Node[K]) Add(val K) {
	node := &Node[K]{Data: val}
	if ll.Next != nil {
		node.Next = ll.Next
	}
	ll.Next = node
}

func (ll *Node[K]) IsEmpty() bool {
	return ll.Next != nil
}

func (ll *Node[K]) Range(f func(data K) bool) {
	head := ll.Next
	for head != nil {
		f(head.Data)
		head = head.Next
	}
}
