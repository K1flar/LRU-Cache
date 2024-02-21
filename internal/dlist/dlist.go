package dlist

type ListNode[K comparable, V any] struct {
	Key   K
	Value V
	Left  *ListNode[K, V]
	Right *ListNode[K, V]
}

func New[K comparable, V any](key K, value V) *ListNode[K, V] {
	return &ListNode[K, V]{
		Key:   key,
		Value: value,
		Left:  nil,
		Right: nil,
	}
}
