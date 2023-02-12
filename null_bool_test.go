package nullany

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNullBool(t *testing.T) {
	assert.True(t, NewNullBool().IsNull())
}

func TestNewNullBoolVal(t *testing.T) {
	assert.False(t, NewNullBoolVal(true).IsNull())
	assert.True(t, NewNullBoolVal(true).Must())
	assert.False(t, NewNullBoolVal(false).IsNull())
	assert.False(t, NewNullBoolVal(false).Must())

}

func TestNullBool_And(t *testing.T) {
	a := NewNullBoolVal(true)
	assert.True(t, a.And(a).Must())
	b := NewNullBoolVal(true)
	assert.True(t, a.And(b).Must())
	c := NewNullBoolVal(false)
	assert.False(t, a.And(c).Must())
	d := NewNullBool()
	assert.True(t, d.IsNull())
	assert.True(t, a.And(d).Must())
	assert.False(t, c.And(d).IsNull())
	assert.False(t, c.And(d).Must())
	assert.False(t, a.And(b).And(c).And(d).Must())
}

func TestNullBool_Get(t *testing.T) {
	a := NewNullBool()
	_, err := a.Get()
	assert.Error(t, err)
	b := NewNullBoolVal(true)
	bv, err := b.Get()
	assert.True(t, bv)
	assert.NoError(t, err)
	c := NewNullBoolVal(false)
	cv, err := c.Get()
	assert.False(t, cv)
	assert.NoError(t, err)
}

func TestNullBool_IsNull(t *testing.T) {
	a := NewNullBool()
	assert.True(t, a.IsNull())
	b := NewNullBoolVal(true)
	assert.False(t, b.IsNull())
	c := NewNullBoolVal(false)
	assert.False(t, c.IsNull())
}

func TestNullBool_Must(t *testing.T) {
	a := NewNullBool()
	assert.False(t, a.Must())
	b := NewNullBoolVal(true)
	assert.True(t, b.Must())
	c := NewNullBoolVal(false)
	assert.False(t, c.Must())
}

func TestNullBool_Or(t *testing.T) {
	a := NewNullBoolVal(true)
	assert.True(t, a.Or(a).Must())
	b := NewNullBoolVal(true)
	assert.True(t, a.Or(b).Must())
	c := NewNullBoolVal(false)
	assert.True(t, a.Or(c).Must())
	d := NewNullBool()
	assert.True(t, a.Or(d).Must())
	assert.False(t, c.Or(d).Must())
	assert.True(t, a.Or(b).Or(c).Or(d).Must())

}
