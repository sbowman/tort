package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// BoolAssertions are tests around boolean values.
type BoolAssertions struct {
	Assertions
	name string
	yesno bool
}

// Bool identifies a bolean variable value and returns test functions for its values.
func (assert Assertions) Bool(value bool) BoolAssertions {
	assert.t.Helper()

	return BoolAssertions{
		Assertions: assert,
		name: "bool",
		yesno: value,
	}
}

// Bool identifies a boolean field on a struct.  If the field isn't present, or isn't a bool,
// generates an error.
func (assert StructAssertions) Bool(field string) BoolAssertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Bool {
		assert.Fatal("field %s is not a bool", name)
	}

	return BoolAssertions{
		Assertions: assert.Assertions,
		name: name,
		yesno: property.Bool(),
	}
}

// Bool looks up an element in a slice expecting it to be a bool.
func (assert SliceAssertions) Bool(idx int) BoolAssertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Bool {
		assert.Fatal("element %d is not a bool", idx)
	}

	return BoolAssertions{
		Assertions: assert.Assertions,
		name: name,
		yesno: property.Bool(),
	}
}

// IsTrue generates an error if the boolean value is false.
func (assert BoolAssertions) IsTrue() {
	assert.t.Helper()

	if assert.yesno == false {
		assert.Failed(`%s is false`, assert.name)
	}
}

// IsFalse generates an error if the boolean value is true.
func (assert BoolAssertions) IsFalse() {
	assert.t.Helper()

	if assert.yesno == true {
		assert.Failed(`%s is true`, assert.name)
	}
}
