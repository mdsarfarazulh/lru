# lru
Sharded mplementation of LRU cache in golang

# Example
```
func main() {
  cache := NewCache(10, 200)
  _, err := cache.Set(K, V)
  item, err := cache.Get(K)
}
```
