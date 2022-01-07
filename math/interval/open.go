package interval

import (
	"fmt"

	"github.com/gtramontina/go-extlib/math/interval/internal"
)

type open[Real internal.Real] struct {
	start, end Real
}

func (open[Real]) seal() (string, string) {
	return "(", ")"
}

func (i open[Real]) List(step Real) []Real {
	list := []Real{}
	for n := i.start + 1; n < i.end; n += step {
		list = append(list, n)
	}
	return list
}

func (i open[Real]) String() string {
	notationStart, notationEnd := i.seal()
	return fmt.Sprintf("Interval%s%v,%v%s", notationStart, i.start, i.end, notationEnd)
}
