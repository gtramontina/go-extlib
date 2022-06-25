package constraints

import "golang.org/x/exp/constraints"

type Real interface {
	constraints.Integer | constraints.Float
}
