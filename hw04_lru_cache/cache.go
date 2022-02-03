package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	node, ok := c.items[key]
	if ok {
		node.Value = cacheItem{key, value}
		c.queue.MoveToFront(node)
	} else {
		c.items[key] = c.queue.PushFront(cacheItem{key, value})
		if c.queue.Len() > c.capacity {
			node := c.queue.Back()
			c.queue.Remove(node)
			delete(c.items, node.Value.(cacheItem).key)
		}
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if node, ok := c.items[key]; ok {
		c.queue.MoveToFront(node)
		return node.Value.(cacheItem).value, ok
	} else {
		return nil, ok
	}
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
