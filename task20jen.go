package main 

import (
	"fmt"
)

type keymap[T any] map[string]T

func (km keymap[T]) set(key string, value T) {
	km[key] = value
}

func (km keymap[T]) get(key string) (T, bool) {
	value, found := km[key]
	return value, found
}

func main() {
	km := make(keymap[int])

	km.set("jack", 10)
	km.set("jill", 20)

	jackValue, jackFound := km.get("jack")

	if jackFound {
		fmt.Printf("jack: %d\n", jackValue)
	}

	jillValue, jillFound := km.get("jill")

	if jillFound {
		fmt.Printf("jill: %d\n", jillValue)
	}
}

