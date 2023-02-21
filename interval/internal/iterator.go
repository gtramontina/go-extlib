package internal

import (
	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/math/constraints"
)

type IntervalContains[Real constraints.Real] interface {
	Contains(Real) bool
}

type Iterator[Real constraints.Real] struct {
	interval IntervalContains[Real]
	step     Real
	current  Real
}

func NewIterator[Real constraints.Real](interval IntervalContains[Real], step Real, current Real) *Iterator[Real] {
	return &Iterator[Real]{
		interval: interval,
		step:     step,
		current:  current,
	}
}

func (i *Iterator[Real]) HasNext() bool {
	return i.interval.Contains(i.current + i.step)
}

func (i *Iterator[Real]) Next() Real {
	if !i.HasNext() {
		panic(iterator.ErrIteratorEmpty)
	}

	i.current += i.step

	return i.current
}

func (i *Iterator[Real]) Collect() []Real {
	collected := []Real{}
	for i.HasNext() {
		collected = append(collected, i.Next())
	}

	return collected
}
