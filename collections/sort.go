package collections

import (
	"sort"

	"github.com/gtramontina/go-extlib/collections/internal/sortable"
	"golang.org/x/exp/constraints"
)

// Sort sorts the given collection by the natural order of its elements in an
// ascending fashion.
func Sort[Type constraints.Ordered](collection []Type) []Type {
	return SortBy(collection, sortable.OrderAscending[Type])
}

// SortDescending sorts the given collection by the natural order of its
// elements in a descending fashion.
func SortDescending[Type constraints.Ordered](collection []Type) []Type {
	return SortBy(collection, sortable.OrderDescending[Type])
}

// SortBy sorts the given collection by the given sorting function. This
// function must return a boolean indicating whether the first argument should
// come before the second argument.
func SortBy[Type any](collection []Type, less func(Type, Type) bool) []Type {
	toBeSorted := make([]Type, len(collection))
	copy(toBeSorted, collection)
	sort.Sort(sortable.New(toBeSorted, less))

	return toBeSorted
}
