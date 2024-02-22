package main

import (
	"fmt"
	"lrucache/pkg/lru"
)

func main() {
	c := lru.New[int, int](10)

	for i := 0; i < 10; i++ {
		c.Set(i, i*10)
	}

	c.Resize(3)

	fmt.Println(c.Keys())
}
