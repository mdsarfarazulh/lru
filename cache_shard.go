package lru

import (
	"errors"
	"fmt"
	"sync"
)

type ShardItem struct {
	Key     string
	Value   interface{}
	LruNode *Node
}

type CacheShard struct {
	index   int
	size    int
	store   map[string]*ShardItem
	mutex   *sync.RWMutex
	lruList *LinkedList
}

type ICacheShard interface {
	Print()
	Get(key string) (*ShardItem, error)
	Set(key string, value interface{}) (*ShardItem, error)
}

func (cs *CacheShard) Print() {
	fmt.Println(cs.store)
}

func (cs *CacheShard) Get(key string) (*ShardItem, error) {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()

	item, ok := cs.store[key]
	if !ok {
		return nil, errors.New("NOT_FOUND")
	}
	node := item.LruNode
	cs.lruList.MoveToTail(node)
	return item, nil
}

func (cs *CacheShard) Set(key string, value interface{}) (*ShardItem, error) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	item := &ShardItem{
		Key:   key,
		Value: value,
	}

	currSize := len(cs.store)
	if currSize+1 >= cs.size {
		cs.lruList.DeleteFromHead()
	}

	node := cs.lruList.InsertAtTail(key)
	item.LruNode = node
	cs.store[key] = item
	return item, nil
}

func NewCacheShard(index, size int) *CacheShard {
	return &CacheShard{
		index:   index,
		size:    size,
		store:   make(map[string]*ShardItem),
		lruList: NewLinkedList(),
		mutex:   &sync.RWMutex{},
	}
}
