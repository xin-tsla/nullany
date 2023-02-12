package nullany

/*
@see https://github.com/xin-tsla/nullany/blob/main/README.md
*/
type NullAny[T any] struct {
	val T
	set bool
}

func NewNullAny[T any](v T) *NullAny[T] {
	na := &NullAny[T]{}
	na.SetVal(v)
	return na
}

func (n *NullAny[T]) SetVal(v T) {
	n.val = v
	n.set = true
}

func (n *NullAny[T]) IsNil() bool {
	return !n.set
}

func (n *NullAny[T]) Do(f func(T)) {
	if n.IsNil() {
		return
	}
	f(n.val)
}

func (n *NullAny[T]) DoThenElse(f1 func(), f2 func(T)) {
	if n.IsNil() {
		f1()
		return
	}
	f2(n.val)
}
