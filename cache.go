package lru

import "github.com/zeebo/xxh3"

type Item struct {
	Key   string
	Value interface{}
}

type ICache interface {
	Get(key string) (*Item, error)
	Set(key string, value interface{}) (*Item, error)
	GetShardIndex(key string) int
	Print()
}

type Cache struct {
	numberOfShards uint64
	shardCapacity  int
	shards         []*CacheShard
}

func (cache *Cache) Print() {
	for i := 0; i < int(cache.numberOfShards); i++ {
		cache.shards[i].Print()
	}
}

func (cache *Cache) GetShardIndex(key string) int {
	return int(xxh3.HashString(key) % cache.numberOfShards)
}

func (cache *Cache) Get(key string) (*Item, error) {
	shardIndex := cache.GetShardIndex(key)
	item, err := cache.shards[shardIndex].Get(key)
	if err != nil {
		return nil, err
	}
	return &Item{
		Key:   item.Key,
		Value: item.Value,
	}, nil
}

func (cache *Cache) Set(key string, value interface{}) (*Item, error) {
	shardIndex := cache.GetShardIndex(key)
	item, err := cache.shards[shardIndex].Set(key, value)
	if err != nil {
		return nil, err
	}
	return &Item{
		Key:   item.Key,
		Value: item.Value,
	}, nil
}

func NewCache(numberOfShards uint64, shardCapacity int) *Cache {
	shards := make([]*CacheShard, int(numberOfShards))
	for i := 0; i < int(numberOfShards); i++ {
		shards[i] = NewCacheShard(i, shardCapacity)
	}
	return &Cache{
		numberOfShards: numberOfShards,
		shardCapacity:  shardCapacity,
		shards:         shards,
	}
}
