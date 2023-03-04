package collections

import "github.com/gtramontina/go-extlib/tuple"

// Zip zips two collections into a collection of tuples of two. The resulting
// collection will be the size of the smaller collection. See also tuple.Of2.
func Zip[A, B any](a []A, b []B) []tuple.OfTwo[A, B] {
	zipped := []tuple.OfTwo[A, B]{}

	for i := 0; i < len(a) && i < len(b); i++ {
		zipped = append(zipped, tuple.Of2(a[i], b[i]))
	}

	return zipped
}
