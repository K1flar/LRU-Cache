package lru

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/K1flar/LRU-Cache/internal/dlist"
)

var (
	ErrElementNotFound = fmt.Errorf("element not found")
	ErrNoFile          = fmt.Errorf("no file")
	ErrLenMore         = fmt.Errorf("length is longer than capacity")
)

const DefaultCapacityCache = 1024

type LRUCache[K comparable, V any] struct {
	len  int
	cap  int
	data map[K]*dlist.ListNode[K, V]
	list *dlist.DoublyLinkedList[K, V]
}

func New[K comparable, V any](cap int) *LRUCache[K, V] {
	if cap <= 0 {
		cap = DefaultCapacityCache
	}

	c := &LRUCache[K, V]{
		len:  0,
		cap:  cap,
		data: make(map[K]*dlist.ListNode[K, V], cap),
		list: dlist.NewDoublyLinkedList[K, V](),
	}

	return c
}

func (c *LRUCache[K, V]) Cap() int {
	return c.cap
}

func (c *LRUCache[K, V]) delete(key K) {
	c.list.RemoveListNode(c.data[key])
	delete(c.data, key)
	c.len--
}

func (c *LRUCache[K, V]) Delete(key K) error {
	if _, ok := c.data[key]; !ok {
		return fmt.Errorf("lru delete: %w", ErrElementNotFound)
	}

	c.delete(key)

	return nil
}

func (c *LRUCache[K, V]) Exist(key K) bool {
	_, ok := c.data[key]
	return ok
}

func (c *LRUCache[K, V]) FlushAll() error {
	c.data = make(map[K]*dlist.ListNode[K, V], c.cap)
	c.list = dlist.NewDoublyLinkedList[K, V]()
	c.len = 0
	return nil
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	if _, ok = c.data[key]; !ok {
		return value, false
	}

	new := dlist.NewListNode[K, V](key, c.data[key].Value)
	c.list.AddToRight(new)
	c.list.RemoveListNode(c.data[key])
	c.data[key] = new

	return c.data[key].Value, true
}

func (c *LRUCache[K, V]) Keys() []K {
	keys := make([]K, 0, c.len)
	for k, _ := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c *LRUCache[K, V]) Len() int {
	return c.len
}

func (c *LRUCache[K, V]) LoadJSON(filePath string) error {
	if len(filePath) == 0 {
		return fmt.Errorf("lru load json: %w", ErrNoFile)
	}

	if filePath[0] != '/' {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("lru load json: %w", err)
		}
		filePath = path.Join(wd, filePath)
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("lru load json: %w", err)
	}

	var loadData map[K]V
	err = json.Unmarshal(b, &loadData)
	if err != nil {
		return fmt.Errorf("testlru load json: %w", err)
	}

	if len(loadData) > c.cap {
		return fmt.Errorf("lru load json: %w", ErrLenMore)
	}

	for k, v := range loadData {
		c.Set(k, v)
	}

	return nil
}

func (c *LRUCache[K, V]) Rename(key, newKey K) error {
	if _, ok := c.data[key]; !ok {
		return fmt.Errorf("lru rename: %w", ErrElementNotFound)
	}

	node := c.data[key]
	node.Key = newKey

	delete(c.data, key)
	c.data[newKey] = node

	return nil
}

func (c *LRUCache[K, V]) Resize(cap int) error {
	if cap <= 0 {
		cap = DefaultCapacityCache
	}

	if cap == c.cap {
		return nil
	}

	newData := make(map[K]*dlist.ListNode[K, V], cap)
	count := c.len

	if cap < c.len {
		count = cap
	}

	node := c.list.Right
	for node != nil && count > 0 {
		newData[node.Key] = node
		count--
		node = node.Left
	}

	if node != nil {
		c.list.Left = node.Right
		c.list.Left.Left = nil
		node.Right = nil
	}

	c.data = newData

	return nil
}

func (c *LRUCache[K, V]) SaveJSON(filePath string) error {
	if len(filePath) == 0 {
		return fmt.Errorf("lru save json: %w", ErrNoFile)
	}

	if filePath[0] != '/' {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("lru save json: %w", err)
		}
		filePath = path.Join(wd, filePath)
	}

	saveData := make(map[K]V, c.cap)
	for k, v := range c.data {
		saveData[k] = v.Value
	}
	b, err := json.Marshal(saveData)
	if err != nil {
		return fmt.Errorf("lru save json: %w", err)
	}

	err = os.WriteFile(filePath, b, 0777)
	if err != nil {
		return fmt.Errorf("lru save json: %w", err)
	}

	return nil
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if _, ok := c.data[key]; ok {
		c.data[key].Value = value
		return
	}

	new := dlist.NewListNode[K, V](key, value)
	c.list.AddToRight(new)
	c.data[key] = new
	c.len++

	if c.len > c.cap {
		d := c.list.Left
		c.list.RemoveListNode(c.list.Left)
		delete(c.data, d.Key)
		c.len--
	}
}

func (c *LRUCache[K, V]) Values() []V {
	values := make([]V, 0, c.len)
	for _, v := range c.data {
		values = append(values, v.Value)
	}
	return values
}
