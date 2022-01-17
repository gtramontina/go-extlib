package constraints

import "constraints"

type Real interface {
	constraints.Integer | constraints.Float
}
