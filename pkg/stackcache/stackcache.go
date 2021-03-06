package stackcache

import "sync"

// Cache describes an cache
type Cache interface {
	Find(id string, createNewItem func(id string) (interface{}, error)) (*interface{}, error)
}

type cache struct {
	items *[]item
	mux   sync.Mutex
}

// an item holds the data and pass also an id
type item struct {
	id   string
	data *interface{}
}

// NewCache creates a new cache
func NewCache(len int) Cache {
	// add dummy-objects to cache
	items := []item{}
	for i := 0; i < len; i++ {
		items = append(items, item{})
	}

	// create cache and return
	return &cache{
		items: &items,
	}
}

// Find finds an item in the stack
func (cache *cache) Find(id string, createNewItem func(id string) (interface{}, error)) (*interface{}, error) {
	var err error
	cache.mux.Lock()
	defer cache.mux.Unlock()

	// check if item is already in cache
	if index := cache.getIndexOfItem(id); index > -1 {
		// item is available => move item to first position
		cache.removeAndAdd(len(*cache.items)-1, (*cache.items)[index])
	} else {
		// create new item and move to first position
		var newItem interface{}
		newItem, err = createNewItem(id)

		// cash only if no error
		// If we add data despite mistakes. We do not know later, whether it is an error or the data is simply empty
		if err == nil {
			cache.removeAndAdd(0, item{id: id, data: &newItem})
		}
	}

	// add data to channel
	return (*cache.items)[len(*cache.items)-1].data, err
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
