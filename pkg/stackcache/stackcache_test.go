package stackcache_test

import (
	"strings"
	"testing"

	"github.com/TobiEiss/stackcache/pkg/stackcache"
)

func TestCache(t *testing.T) {
	createNewItem := func(id string) (interface{}, error) {
		return strings.Replace(id, "ID", "Data", 1), nil
	}

	// Create a new cache
	stack := stackcache.NewCache(5)

	// find data
	dataInterface, err := stack.Find("myID1", createNewItem)
	data1 := (*(dataInterface)).(string)
	if data1 != "myData1" && err != nil {
		t.Fail()
	}
}
