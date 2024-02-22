package dlist

type ListNode[K comparable, V any] struct {
	Key   K
	Value V
	Left  *ListNode[K, V]
	Right *ListNode[K, V]
}

func NewListNode[K comparable, V any](key K, value V) *ListNode[K, V] {
	return &ListNode[K, V]{
		Key:   key,
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

type DoublyLinkedList[K comparable, V any] struct {
	Left  *ListNode[K, V]
	Right *ListNode[K, V]
}

// func NewDoublyLinkedList[K comparable, V any]() *DoublyLinkedList[K, V] {
// 	var dk K
// 	var dv V
// 	l := NewListNode[K, V](dk, dv)
// 	r := NewListNode[K, V](dk, dv)
// 	l.Left, r.Right = nil, nil
// 	l.Right, r.Left = r, l
// 	return &DoublyLinkedList[K, V]{
// 		Left:  l,
// 		Right: r,
// 	}
// }

func NewDoublyLinkedList[K comparable, V any]() *DoublyLinkedList[K, V] {
	return &DoublyLinkedList[K, V]{
		Left:  nil,
		Right: nil,
	}
}

func (l *DoublyLinkedList[K, V]) AddToRight(node *ListNode[K, V]) {
	if node == nil {
		return
	}

	if l.Right == nil {
		l.Right = node
		l.Right.Left = nil
		l.Right.Right = nil
		return
	}

	node.Right = nil
	node.Left = l.Right

	if l.Left == nil {
		l.Left = l.Right
		l.Left.Left = nil
	}

	l.Right.Right = node
	l.Right = node
}

func (l *DoublyLinkedList[K, V]) RemoveListNode(node *ListNode[K, V]) {
	if node == nil {
		return
	}

	if node.Right == nil {
		if node.Left != nil {
			node.Left.Right = nil
		}
		l.Right = node.Left
		return
	}

	if node.Left == nil {
		node.Right.Left = nil
		l.Left = node.Right
		return
	}

	node.Left.Right = node.Right
	node.Right.Left = node.Left
}
