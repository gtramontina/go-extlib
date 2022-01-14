package sortable

type Sortable[Type any] struct {
	collection []Type
	less       func(i, j Type) bool
}

func New[Type any](collection []Type, less func(i, j Type) bool) Sortable[Type] {
	return Sortable[Type]{
		collection: collection,
		less:       less,
	}
}

func (s Sortable[Type]) Len() int {
	return len(s.collection)
}

func (s Sortable[Type]) Swap(i, j int) {
	s.collection[i], s.collection[j] = s.collection[j], s.collection[i]
}

func (s Sortable[Type]) Less(i, j int) bool {
	return s.less(s.collection[i], s.collection[j])
}
