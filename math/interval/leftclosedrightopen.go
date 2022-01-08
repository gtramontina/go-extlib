package interval

import (
	"fmt"

	"github.com/gtramontina/go-extlib/math/interval/internal"
)

type leftclosedrightopen[Real internal.Real] struct {
	start, end Real
}

func (leftclosedrightopen[Real]) seal() (string, string) {
	return "[", ")"
}

func (i leftclosedrightopen[Real]) List(step Real) []Real {
	list := []Real{}
	for n := i.start; n < i.end; n += step {
		list = append(list, n)
	}
	return list
}

func (i leftclosedrightopen[Real]) Contains(n Real) bool {
	return n >= i.start && n < i.end
}

func (i leftclosedrightopen[Real]) String() string {
	notationStart, notationEnd := i.seal()
	return fmt.Sprintf("Interval%s%v,%v%s", notationStart, i.start, i.end, notationEnd)
}
