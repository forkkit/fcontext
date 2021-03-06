package fcontext

import (
	"testing"

	"github.com/posener/contexttest"
	"github.com/stretchr/testify/assert"
)

func TestWithValue(t *testing.T) {
	contexttest.TestWithValue(WithValue)(t)
}

func TestWithValues(t *testing.T) {
	contexttest.TestWithValue(func(ctx Context, k, v interface{}) Context {
		return WithValues(ctx, k, v)
	})(t)
}
func TestWithValues_multipleValues(t *testing.T) {
	t.Parallel()
	ctx := WithValues(Background(), 0, 0, 1, 1)
	assert.Equal(t, 0, ctx.Value(0))
	assert.Equal(t, 1, ctx.Value(1))
}

func TestWithValues_panic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		WithValues(Background(), 0)
	})
}
