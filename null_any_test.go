package nullany

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullAny_Do(t *testing.T) {
	na := NewNullAny[int](1)
	assert.False(t, na.IsNil())
	cnt := 0
	na.Do(
		func(i int) {
			cnt = i
		},
	)
	assert.Equal(t, 1, cnt)
}

func TestNullAny_IsNil(t *testing.T) {
	na := NullAny[bool]{}
	assert.True(t, na.IsNil())
}

func TestNullAny_SetVal(t *testing.T) {
	na := NullAny[bool]{}
	assert.True(t, na.IsNil())
	na.SetVal(true)
	assert.False(t, na.IsNil())
}

func TestNullAny_NewNullAny(t *testing.T) {
	na := NewNullAny[int](1)
	assert.False(t, na.IsNil())
}
