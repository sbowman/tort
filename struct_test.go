package tort_test

import (
	"testing"

	"github.com/sbowman/tort"
)

func TestStructNil(t *testing.T) {
	assert := tort.For(t)

	type Sample struct {
	}

	var ptr *Sample
	assert.Struct(ptr).IsNil()

	ptr = new(Sample)
	assert.Struct(ptr).IsNotNil()

	var sample Sample
	assert.Struct(sample).IsNotNil()
}
