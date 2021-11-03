package homework_4

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type LruCache struct {
	size     int
	elements ForwardList
	mp       map[string]*Node
}

func NewCache(size int) Cache {
	return &LruCache{size: size, elements: NewList(), mp: make(map[string]*Node)}
}

func (c *LruCache) Set(key string, value interface{}) (ok bool) {
	node, ok := c.mp[key]
	if ok {
		node.Value = value
		c.elements.MoveToFront(node)
	} else {
		node = c.elements.PushFront(value)
		c.mp[key] = node
	}
	return
}

func (c *LruCache) Get(key string) (interface{}, bool) {
	_, ok := c.mp[key]
	if ok {
		return c.mp[key].Value, true
	}
	return nil, false
}

func (c *LruCache) Clear() {
	// @TBD
}
