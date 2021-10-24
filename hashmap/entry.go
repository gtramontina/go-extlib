package hashmap

// Entry holds a key/value pair. Use Pair to construct a pair.
type Entry[Key any, Value any] struct {
	key   Key
	value Value
}

// Pair builds a new Entry holding the given key/value pair.
func Pair[Key any, Value any](key Key, value Value) Entry[Key, Value] {
	return Entry[Key, Value]{key, value}
}
