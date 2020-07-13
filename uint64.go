package tort

import (
	"fmt"
	"reflect"
)

// Code generated from templates/uints.tmpl; DO NOT EDIT.

// Uint64Assertions are tests around integer (uint64) values.
type Uint64Assertions struct {
	Assertions
	name string
	num  uint64
}

// Uint64 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Uint64(value uint64) Uint64Assertions {
	return Uint64Assertions{
		Assertions: assert,
		name:       "uint64",
		num:        value,
	}
}

// Uint looks for the given struct field, confirms it's an uint64, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Uint64(field string) Uint64Assertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Uint64 {
		assert.Fatal("field %s is not an uint64", name)
	}

	return Uint64Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint64(property.Uint()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Uint64Assertions) Equals(other uint64) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Uint64Assertions) NotEquals(other uint64) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Uint64Assertions) GreaterThan(other uint64) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Uint64Assertions) LessThan(other uint64) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
