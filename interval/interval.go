package interval

import (
	"fmt"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/math/constraints"
)

// Interval is a set of real numbers that contains all real numbers lying
// between any two numbers of the set.
//
//	Notation:
//	• Open:                (𝑎,𝑏) = { 𝑥 ∈ ℝ | 𝑎 < 𝑥 < 𝑏 }
//	• LeftClosedRightOpen: [𝑎,𝑏) = { 𝑥 ∈ ℝ | 𝑎 ≤ 𝑥 < 𝑏 }
//	• LeftOpenRightClosed: (𝑎,𝑏] = { 𝑥 ∈ ℝ | 𝑎 < 𝑥 ≤ 𝑏 }
//	• Closed:              [𝑎,𝑏] = { 𝑥 ∈ ℝ | 𝑎 ≤ 𝑥 ≤ 𝑏 }
//
//	Graphical Representation:
//	• Open:                𝑎 ○————○ 𝑏
//	• LeftClosedRightOpen: 𝑎 ●————○ 𝑏
//	• LeftOpenRightClosed: 𝑎 ○————● 𝑏
//	• Closed:              𝑎 ●————● 𝑏
type Interval[Real constraints.Real] interface {
	fmt.Stringer

	// seal is used internally as a way of limiting external implementations.
	seal() (string, string)

	// List represents this Interval as an ordered array, from start to end,
	// containing all numbers distanced by a given step size.
	List(Real) []Real

	// Contains checks if this interval contains the given number.
	Contains(Real) bool

	// Iterator returns an iterator that can be used to iterate over all numbers
	// in this interval.
	Iterator(Real) iterator.Iterator[Real]
}

// Open creates an open interval, where both start and end are excluded.
//   - Graphical: 𝑎 ○————○ 𝑏
//   - Notation:  (𝑎,𝑏) = { 𝑥 ∈ ℝ | 𝑎 < 𝑥 < 𝑏 }
//
// See also: LeftClosedRightOpen, LeftOpenRightClosed, Closed.
func Open[Real constraints.Real](start, end Real) Interval[Real] {
	return open[Real]{start: start, end: end}
}

// LeftClosedRightOpen creates a left-closed, right-open interval, where start
// is included and end is excluded.
//   - Graphical: 𝑎 ●————○ 𝑏
//   - Notation:  [𝑎,𝑏) = { 𝑥 ∈ ℝ | 𝑎 ≤ 𝑥 < 𝑏 }
//
// See also: Open, LeftOpenRightClosed, Closed.
func LeftClosedRightOpen[Real constraints.Real](start, end Real) Interval[Real] {
	return leftclosedrightopen[Real]{start: start, end: end}
}

// LeftOpenRightClosed creates a left-open, right-closed interval, where start
// is excluded and end is included.
//   - Graphical: 𝑎 ○————● 𝑏
//   - Notation:  (𝑎,𝑏] = { 𝑥 ∈ ℝ | 𝑎 < 𝑥 ≤ 𝑏 }
//
// See also: Open, LeftClosedRightOpen, Closed.
func LeftOpenRightClosed[Real constraints.Real](start, end Real) Interval[Real] {
	return leftopenrightclosed[Real]{start: start, end: end}
}

// Closed creates a closed interval, where both start and end are included.
//   - Graphical: 𝑎 ●————● 𝑏
//   - Notation:  [𝑎,𝑏] = { 𝑥 ∈ ℝ | 𝑎 ≤ 𝑥 ≤ 𝑏 }
//
// See also: Open, LeftClosedRightOpen, LeftOpenRightClosed.
func Closed[Real constraints.Real](start, end Real) Interval[Real] {
	return closed[Real]{start: start, end: end}
}
