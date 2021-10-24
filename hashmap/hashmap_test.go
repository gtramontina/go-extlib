package hashmap_test

import (
	"github.com/gtramontina/go-collections/hashmap"
	"github.com/gtramontina/go-collections/internal/assert"
	"github.com/gtramontina/go-collections/set"
	"testing"
)

func TestHashMap(t *testing.T) {
	t.Run("has size", func(t *testing.T) {
		assert.Eq(t, hashmap.New[string, int]().Size(), 0)
		assert.Eq(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Size(), 1)
		assert.Eq(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Size(), 2)
	})

	t.Run("tells whether it is empty", func(t *testing.T) {
		assert.True(t, hashmap.New[string, int]().Empty())
		assert.False(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Empty())
		assert.False(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Empty())
	})

	t.Run("can be compared to other hash maps", func(t *testing.T) {
		assert.Equals(t, hashmap.New[string, int](), hashmap.New[string, int]())
		assert.NotEquals(t, hashmap.New[string, int](), hashmap.New[string, int](hashmap.Pair("key1", 1)))
		assert.Equals(t, hashmap.New[string, int](hashmap.Pair("key1", 1)), hashmap.New[string, int](hashmap.Pair("key1", 1)))
		assert.Equals(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)), hashmap.New[string, int](hashmap.Pair("key2", 2), hashmap.Pair("key1", 1)))
	})

	t.Run("tells whether it has a value for the given key", func(t *testing.T) {
		assert.False(t, hashmap.New[string, int]().HasKey("dummy"))
		filledMap := hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2))
		assert.True(t, filledMap.HasKey("key1"))
		assert.True(t, filledMap.HasKey("key2"))
		assert.False(t, filledMap.HasKey("nope"))
	})

	t.Run("allows putting new entries", func(t *testing.T) {
		assert.Equals(t, hashmap.New[string, int]().Put("key1", 1), hashmap.New[string, int](hashmap.Pair("key1", 1)))
		assert.Equals(t, hashmap.New[string, int]().Put("key1", 1).Put("key2", 2), hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)))

		t.Run("replaces value when same key", func(t *testing.T) {
			assert.Equals(t, hashmap.New[string, int]().Put("key1", 1).Put("key1", 0), hashmap.New[string, int](hashmap.Pair("key1", 0)))
		})
	})

	t.Run("allows removing entries", func(t *testing.T) {
		assert.Equals(t, hashmap.New[string, int]().Remove("dummy"), hashmap.New[string, int]())
		assert.Equals(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Remove("key1"), hashmap.New[string, int]())
		assert.Equals(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Remove("key1"), hashmap.New[string, int](hashmap.Pair("key2", 2)))
		assert.Equals(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Remove("nope"), hashmap.New[string, int](hashmap.Pair("key1", 1)))
	})

	t.Run("can retrieve values based on keys", func(t *testing.T) {
		assert.Eq(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Get("key1"), 1)
		assert.Eq(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Get("key2"), 2)
		assert.Eq(t, hashmap.New[string, int](hashmap.Pair("key0", 0), hashmap.Pair("key3", 3)).Get("key3"), 3)

		t.Run("panics when no value is associated with given key", func(t *testing.T) {
			assert.Panic(t, func() { hashmap.New[string, int]().Get("unknown") }, "hashmap: key not found")
		})
	})

	t.Run("allows accessing all keys", func(t *testing.T) {
		assert.DeepEqual(t, hashmap.New[string, int]().Keys(), set.New[string]())
		assert.DeepEqual(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Keys(), set.New("key1"))
		assert.DeepEqual(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Keys(), set.New("key1", "key2"))
	})

	t.Run("allows accessing all values", func(t *testing.T) {
		assert.DeepEqual(t, hashmap.New[string, int]().Values(), set.New[int]())
		assert.DeepEqual(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Values(), set.New(1))
		assert.DeepEqual(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Values(), set.New(1, 2))
	})

	t.Run("allows accessing all entries", func(t *testing.T) {
		assert.DeepEqual(t, hashmap.New[string, int]().Entries(), set.New[hashmap.Entry[string, int]]())
		assert.DeepEqual(t, hashmap.New[string, int](hashmap.Pair("key1", 1)).Entries(), set.New[hashmap.Entry[string, int]](hashmap.Pair("key1", 1)))
		assert.DeepEqual(t, hashmap.New[string, int](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)).Entries(), set.New[hashmap.Entry[string, int]](hashmap.Pair("key1", 1), hashmap.Pair("key2", 2)))
	})
}
