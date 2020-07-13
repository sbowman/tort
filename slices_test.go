package tort_test

import (
	"testing"

	"github.com/sbowman/tort"
)

func TestSlices(t *testing.T) {
	assert := tort.For(t)

	slice := []int{3, 1, 4, 1, 5, 9, 2}
	assert.Slice(slice).Length(7)
	assert.Slice(slice).MoreThan(6)
	assert.Slice(slice).FewerThan(8)
	assert.Slice(slice).Int(4).Equals(5)
}
