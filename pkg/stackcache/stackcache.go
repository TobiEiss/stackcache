package stackcache

// Cache describes an cache
type Cache interface {
	Find(id string) chan interface{}
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

func (cache *cache) Find(id string) chan interface{} {
	channel := make(chan interface{})
	go func(c chan interface{}) {
		c <- cache.getNewItem(id)
	}(channel)
	return channel
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
