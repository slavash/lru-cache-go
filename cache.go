package cache

import (
	"container/list"
)

type item struct {
	k string
	v interface{}
}

type cache struct {
	cache map[string]*list.Element
	order *list.List
	limit int
}

func New(size int) *cache {
	return new(cache).init(size)
}

func (c *cache) init(n int) *cache {
	return &cache{
		cache: make(map[string]*list.Element),
		order: list.New(),
		limit: n,
	}
}

func (c *cache) Set(key string, val interface{}) bool {

	// existing key, just renew the value
	e, exists := c.cache[key]
	if exists {
		e.Value = item{key, val}
		c.order.MoveToFront(e) // I consider the Set as a usage as well, so item is moved at the beginning

		return true
	}

	// new key, no overflow, item placed at the beginning
	if len(c.cache) < c.limit {
		c.cache[key] = c.order.PushFront(item{key, val})

		return true
	}

	// new key, overflow:
	// - take latest item key
	e = c.order.Back()
	i, casted := e.Value.(item)
	if !casted {
		return false
	}

	// - delete it from the map
	delete(c.cache, i.k)
	// - delete it from the queue
	c.order.Remove(e)

	// place new item at the beginning and add it to the cache map
	newItem := c.order.PushFront(item{key, val})
	c.cache[key] = newItem

	return true
}

func (c *cache) Get(key string) (interface{}, bool) {

	val, ok := c.cache[key]
	if !ok {
		return "", false
	}

	c.order.MoveToFront(val)
	i := val.Value.(item)

	return i.v, true
}
