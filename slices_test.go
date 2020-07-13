package tort_test

import (
	"testing"

	"github.com/sbowman/tort"
)

func TestBoolSlices(t *testing.T) {
	assert := tort.For(t)

	slice := []bool{true, true, false}
	assert.Slice(slice).Length(3)
	assert.Slice(slice).Bool(0).IsTrue()
	assert.Slice(slice).Bool(2).IsFalse()
}

func TestIntSlices(t *testing.T) {
	assert := tort.For(t)

	slice := []int{3, 1, 4, 1, 5, 9, 2}
	assert.Slice(slice).Length(7)
	assert.Slice(slice).MoreThan(6)
	assert.Slice(slice).FewerThan(8)
	assert.Slice(slice).Int(4).Equals(5)
}
