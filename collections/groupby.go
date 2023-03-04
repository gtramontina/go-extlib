package collections

import (
	"github.com/gtramontina/go-extlib/hashmap"
)

// GroupBy groups a slice of items by a given key function.
func GroupBy[Type any, Key any](slice []Type, keyFunc func(Type) Key) hashmap.HashMap[Key, []Type] {
	grouped := hashmap.New[Key, []Type]()

	for _, item := range slice {
		key := keyFunc(item)
		group := append(grouped.MaybeGet(key).UnwrapOrElse(func() []Type { return []Type{} }), item)
		grouped = grouped.Put(key, group)
	}

	return grouped
}
