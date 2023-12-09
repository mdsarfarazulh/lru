package lru

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestGetSet(t *testing.T) {
	cache := NewCache(NewCacheConfig(10, 20000))
	const count int = 10000 //00000
	var keys string
	var values string

	for i := 0; i < count; i++ {
		keys = uuid.NewString()
		values = uuid.NewString()

		_, err := cache.Set(keys, values)
		if err != nil {
			fmt.Printf("Failed to set (K, V): (%s, %s) due to (error: %s)\n", keys, values, err)
		}

		item, err := cache.Get(keys)
		if err != nil {
			fmt.Printf("Failed to get (K, V): (%s, %s) due to (error: %s)\n", keys, values, err)
		} else {
			fmt.Printf("%d - (K, V) -> (%s, %s)\n", i, item.Key, item.Value)
		}
	}

}
