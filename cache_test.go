package lru

import (
	"encoding/json"
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

		value := Item{Key: keys, Value: values}
		_, err := cache.Set(keys, value)
		if err != nil {
			fmt.Printf("Failed to set (K, V): (%s, %s) due to (error: %s)\n", keys, values, err)
		}

		item, err := cache.Get(keys)
		if err != nil {
			fmt.Printf("Failed to get (K, V): (%s, %s) due to (error: %s)\n", keys, values, err)
		} else {
			b, err := json.MarshalIndent(item, "", "  ")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(b))
		}
	}
}
