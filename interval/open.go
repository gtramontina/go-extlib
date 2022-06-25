package interval

import (
	"fmt"

	"github.com/gtramontina/go-extlib/math/constraints"
)

type open[Real constraints.Real] struct {
	start, end Real
}

func (open[Real]) seal() (string, string) {
	return "(", ")"
}

func (i open[Real]) List(step Real) []Real {
	list := []Real{}
	for n := i.start + step; n < i.end; n += step {
		list = append(list, n)
	}

	return list
}

func (i open[Real]) Contains(n Real) bool {
	return n > i.start && n < i.end
}

func (i open[Real]) String() string {
	notationStart, notationEnd := i.seal()

	return fmt.Sprintf("Interval%s%v,%v%s", notationStart, i.start, i.end, notationEnd)
}
