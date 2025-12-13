package heap

type Heap[T any] struct {
	elements []T
	less     func(i, j T) bool
}

func (h Heap[T]) Len() int           { return len(h.elements) }
func (h Heap[T]) Less(i, j int) bool { return h.less(h.elements[i], h.elements[j]) }
func (h Heap[T]) Swap(i, j int)      { h.elements[i], h.elements[j] = h.elements[j], h.elements[i] }

func (h *Heap[T]) Push(x interface{}) {
	h.elements = append(h.elements, x.(T))
}

func (h *Heap[T]) Pop() interface{} {
	old := h.elements
	n := len(old)
	x := old[n-1]
	h.elements = old[0 : n-1]
	return x
}

func New[T any](less func(i, j T) bool) *Heap[T] {
	return &Heap[T]{elements: make([]T, 0), less: less}
}
