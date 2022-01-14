package sortable

import (
	"constraints"
	"math"
)

func OrderAscending[Type constraints.Ordered](i Type, j Type) bool {
	if iFloat, ok := any(i).(float64); ok {
		return i < j || (math.IsNaN(iFloat) && !math.IsNaN(any(j).(float64)))
	}
	return i < j
}

func OrderDescending[Type constraints.Ordered](i Type, j Type) bool {
	if iFloat, ok := any(i).(float64); ok {
		return i > j || (math.IsNaN(any(j).(float64)) && !math.IsNaN(iFloat))
	}
	return i > j
}
