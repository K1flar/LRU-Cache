package main

import (
	"fmt"
	"lrucache/pkg/lru"
)

func main() {
	c := lru.New[string, int](5)

	c.Set("0", 10)
	c.Set("1", 20)
	c.Set("2", 30)
	c.Set("3", 40)

	fmt.Println(c.Get("0"))
	fmt.Println(c.Get("1"))
	fmt.Println(c.Get("2"))
	fmt.Println(c.Get("3"))
}
