package stackcache

// Cache describes an cache
type Cache interface {
	Find(id string) interface{}
}

type cache struct {
	items      []item
	getNewItem func(id string) interface{}
}

// an item holds the data and pass also an id
type item struct {
	id   string
	data *interface{}
}

// NewCache creates a new cache
func NewCache(len int, getNewItem func(id string) interface{}) Cache {
	return &cache{
		items:      []item{},
		getNewItem: getNewItem,
	}
}

func (cache *cache) Find(id string) interface{} {
	return cache.getNewItem(id)
}

// getIndexOfItem search for an Item
func (cache *cache) getIndexOfItem(id string) int {
	for index, item := range cache.items {
		if item.id == id {
			return index
		}
	}
	return -1
}
