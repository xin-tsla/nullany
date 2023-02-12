package nullany

type NullAny[T any] struct {
	val T
}

func (n *NullAny[T]) SetVal(v T) {
	n.val = v
}

func (n *NullAny[T]) IsNil() bool {
	return n.val == nil
}

func (n *NullAny[T]) Do(f func(T)) {
	if n.IsNil() {
		return
	}
	f(n.val)
}
