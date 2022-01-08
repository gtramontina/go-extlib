package internal

import "constraints"

type Real interface {
	constraints.Integer | constraints.Float
}
