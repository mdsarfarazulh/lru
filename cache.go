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
	Info()
}

type Cache struct {
	cacheConfig CacheConfig
	shards      []*CacheShard
}

func (cache *Cache) Info() {
	for i := 0; i < int(cache.cacheConfig.NumberOfShards); i++ {
		cache.shards[i].Info()
	}
}

func (cache *Cache) Print() {
	for i := 0; i < int(cache.cacheConfig.NumberOfShards); i++ {
		cache.shards[i].Print()
	}
}

func (cache *Cache) GetShardIndex(key string) int {
	return int(xxh3.HashString(key) % cache.cacheConfig.NumberOfShards)
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

func NewCache(cacheConfig CacheConfig) *Cache {
	shards := make([]*CacheShard, int(cacheConfig.NumberOfShards))
	for i := 0; i < int(cacheConfig.NumberOfShards); i++ {
		shards[i] = NewCacheShard(i, cacheConfig.ShardCapacity)
	}
	return &Cache{
		cacheConfig: cacheConfig,
		shards:      shards,
	}
}
