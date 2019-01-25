package stackcache

import "sync"

// Cache describes an cache
type Cache interface {
	Find(id string) chan *interface{}
}

type cache struct {
	items      *[]item
	getNewItem func(id string) interface{}
	mux        sync.Mutex
}

// an item holds the data and pass also an id
type item struct {
	id   string
	data *interface{}
}

// NewCache creates a new cache
func NewCache(len int, getNewItem func(id string) interface{}) Cache {
	// add dummy-objects to cache
	items := []item{}
	for i := 0; i < len; i++ {
		items = append(items, item{})
	}

	// create cache and return
	return &cache{
		items:      &items,
		getNewItem: getNewItem,
	}
}

// Find finds an item in the stack
func (cache *cache) Find(id string) chan *interface{} {
	channel := make(chan *interface{})
	go func(c chan *interface{}) {
		defer close(c)
		cache.mux.Lock()
		defer cache.mux.Unlock()

		// check if item is already in cache
		if index := cache.getIndexOfItem(id); index > -1 {
			// item is available => move item to first position
			cache.removeAndAdd(len(*cache.items), (*cache.items)[index])
		} else {
			// create new item and move to first position
			data := cache.getNewItem(id)
			cache.removeAndAdd(0, item{id: id, data: &data})
		}

		// add data to channel
		c <- (*cache.items)[len(*cache.items)-1].data
	}(channel)
	return channel
}

// getIndexOfItem search for an Item
func (cache *cache) getIndexOfItem(id string) int {
	for index, item := range *cache.items {
		if item.id == id {
			return index
		}
	}
	return -1
}

// Remove on item by index and add a new one
// first item = len(items)
// last item = 0
func (cache *cache) removeAndAdd(old int, new item) {
	items := append(append((*cache.items)[:old], (*cache.items)[old+1:]...), new)
	cache.items = &items
}
