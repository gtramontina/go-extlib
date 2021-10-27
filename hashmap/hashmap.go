package hashmap

import (
	"github.com/gtramontina/go-collections/internal/hash"
	"github.com/gtramontina/go-collections/set"
	"reflect"
)

// HashMap implements a structure that maps keys to values. It uses a hash
// function (hash.Calc) to compute a hash code used to identify values.
type HashMap[Key any, Value any] struct {
	entries map[uint64]Entry[Key, Value]
}

// New creates a new HashMap containing the given key/value entry pairs.
func New[Key any, Value any](entries ...Entry[Key, Value]) HashMap[Key, Value] {
	newEntries := make(map[uint64]Entry[Key, Value], len(entries))
	for _, entry := range entries {
		newEntries[hash.Calc(entry.key)] = entry
	}

	return HashMap[Key, Value]{entries: newEntries}
}

// Size returns the number of key/value pair entries this HashMap holds.
func (m HashMap[Key, Value]) Size() int {
	return len(m.entries)
}

// Empty returns true when this HashMap does not hold any entries; false
// otherwise.
func (m HashMap[Key, Value]) Empty() bool {
	return m.Size() == 0
}

// Put creates a new HashMap containing all existing entries plus the newly
// given key/value pair.
// Note: if the calculated hash for the given Key collides with an existing
// entry, the value will be replaced with the given Value.
func (m HashMap[Key, Value]) Put(key Key, value Value) HashMap[Key, Value] {
	newEntries := make(map[uint64]Entry[Key, Value], len(m.entries)+1)
	for h, entry := range m.entries {
		newEntries[h] = entry
	}

	newEntries[hash.Calc(key)] = Pair[Key, Value](key, value)

	return HashMap[Key, Value]{entries: newEntries}
}

// Remove creates a new HashMap containing all existing entries but the one
// whose hash code matches the given Key's. If none is found, the resulting
// HashMap is equal to the original.
func (m HashMap[Key, Value]) Remove(key Key) HashMap[Key, Value] {
	newEntries := make(map[uint64]Entry[Key, Value], len(m.entries)+1)
	hashedKey := hash.Calc(key)
	for h, entry := range m.entries {
		if h != hashedKey {
			newEntries[h] = entry
		}
	}

	return HashMap[Key, Value]{entries: newEntries}
}

// MustGet retrieves the Value for the given Key. Panics when key is not found.
// TODO: implement Maybe[Value]
func (m HashMap[Key, Value]) MustGet(key Key) Value {
	entry, ok := m.entries[hash.Calc(key)]
	if !ok {
		panic("hashmap: key not found")
	}

	return entry.value
}

// Keys returns a set.Set of all keys contained in this HashMap.
func (m HashMap[Key, Value]) Keys() set.Set[Key] {
	keys := set.New[Key]()
	for _, entry := range m.entries {
		keys = keys.Add(entry.key)
	}

	return keys
}

// Values returns a set.Set of all values contained in this HashMap.
func (m HashMap[Key, Value]) Values() set.Set[Value] {
	values := set.New[Value]()
	for _, entry := range m.entries {
		values = values.Add(entry.value)
	}

	return values
}

// Entries returns a set.Set of all key/value pair entries contained in this
// HashMap.
func (m HashMap[Key, Value]) Entries() set.Set[Entry[Key, Value]] {
	entries := set.New[Entry[Key, Value]]()
	for _, entry := range m.entries {
		entries = entries.Add(entry)
	}

	return entries
}

// HasKey returns true if this HashMap contains a Value for the given Key; false
// otherwise.
func (m HashMap[Key, Value]) HasKey(key Key) bool {
	_, has := m.entries[hash.Calc(key)]

	return has
}

// Equals compares this HashMap with another HashMap. Returns true when all keys
// and values are the same; false otherwise.
func (m HashMap[Key, Value]) Equals(other HashMap[Key, Value]) bool {
	return reflect.DeepEqual(m, other)
}
