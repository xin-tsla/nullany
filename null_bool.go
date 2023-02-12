package nullany

import "errors"

type NullBool uint8

const (
	Null NullBool = iota
	False
	True
)

func NewNullBool() NullBool {
	return Null
}

func NewNullBoolVal(b bool) NullBool {
	if b {
		return True
	}
	return False
}

func (n NullBool) IsNull() bool {
	return n == Null
}

func (n NullBool) Get() (bool, error) {
	switch n {
	case False:
		return false, nil
	case True:
		return true, nil
	default:
		return false, errors.New("null")
	}
}

func (n NullBool) Must() bool {
	if n == True {
		return true
	}
	return false
}

type op uint8

const (
	and op = iota
	or
)

func handleBinaryOp(a, b NullBool, o op) NullBool {
	if a == Null && b == Null {
		return Null
	}
	if a == Null {
		return b
	}
	if b == Null {
		return a
	}
	l := a.Must()
	r := b.Must()
	switch o {
	case and:
		return NewNullBoolVal(l && r)
	case or:
		return NewNullBoolVal(l || r)
	}
	return Null
}

func (n NullBool) And(nb NullBool) NullBool {
	return handleBinaryOp(n, nb, and)
}

func (n NullBool) Or(nb NullBool) NullBool {
	return handleBinaryOp(n, nb, or)
}
