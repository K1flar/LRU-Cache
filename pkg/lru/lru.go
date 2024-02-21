package lru

type listNode[K comparable, V any] struct {
	Key   K
	Value V
	Left  *listNode[K, V]
	Right *listNode[K, V]
}

func newListNode[K comparable, V any](key K, value V) *listNode[K, V] {
	return &listNode[K, V]{
		Key:   key,
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

type LRUCache[K comparable, V any] struct {
	len   int
	cap   int
	data  map[K]*listNode[K, V]
	left  *listNode[K, V]
	right *listNode[K, V]
}

func New[K comparable, V any](cap int) *LRUCache[K, V] {
	c := &LRUCache[K, V]{
		len:   0,
		cap:   cap,
		data:  make(map[K]*listNode[K, V], cap),
		left:  nil,
		right: nil,
	}

	return c
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	var res V
	if _, exists := c.data[key]; !exists {
		return res, false
	}

	new := newListNode[K, V](c.data[key].Key, c.data[key].Value)
	c.addListNode(new)
	c.removeListNode(c.data[key])
	c.data[key] = new

	return c.data[key].Value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if _, exists := c.data[key]; exists {
		c.data[key].Value = value
		return
	}

	new := newListNode[K, V](key, value)
	c.addListNode(new)
	c.data[key] = new
	c.len++

	if c.len > c.cap {
		d := c.left
		c.removeListNode(c.left)
		delete(c.data, d.Key)
		c.left = d.Right
		c.len--
	}
}

func (c *LRUCache[K, V]) addListNode(node *listNode[K, V]) {
	if c.right == nil {
		c.right = node
		return
	}

	c.right.Right = node
	node.Right = nil
	node.Left = c.right
	if c.left == nil {
		c.left = c.right
		c.left.Left = nil
	}
	c.right = node
}

func (c *LRUCache[K, V]) removeListNode(node *listNode[K, V]) {
	if node.Left == nil {
		c.left = node.Right
		c.left.Left = nil
		return
	}

	if node.Right == nil {
		c.right = node.Left
		c.right.Right = nil
		return
	}

	node.Left.Right = node.Right
	node.Right.Left = node.Left
}
