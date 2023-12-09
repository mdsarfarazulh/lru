package lru

type CacheConfig struct {
	NumberOfShards uint64
	ShardCapacity  int
}

func NewCacheConfig(numberOfShards uint64, shardCapacity int) CacheConfig {
	return CacheConfig{
		NumberOfShards: numberOfShards,
		ShardCapacity:  shardCapacity,
	}
}
