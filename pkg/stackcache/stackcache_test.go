package stackcache_test

import (
	"strings"
	"testing"

	"github.com/TobiEiss/stackcache/pkg/stackcache"
)

func TestCache(t *testing.T) {
	// Create a new cache
	stack := stackcache.NewCache(5, func(id string) interface{} {
		return strings.Replace(id, "ID", "Data", 1)
	})

	// find data
	data1 := (<-stack.Find("myID1")).(string)
	if data1 != "myData1" {
		t.Fail()
	}
}
