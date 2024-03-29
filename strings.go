package tort

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// StringAssertions are tests around string values.
type StringAssertions struct {
	Assertions
	name string
	str  string
}

// String identifies a string variable value and returns test functions for its values.
func (assert Assertions) String(value string) StringAssertions {
	assert.t.Helper()

	return StringAssertions{
		Assertions: assert,
		name:       "string",
		str:        value,
	}
}

// String identifies a string field on a struct.  If the field isn't present, or isn't a string,
// generates an error.
func (assert StructAssertions) String(field string) StringAssertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.String {
		assert.Fatal("field %s is not a string", name)
	}

	return StringAssertions{
		Assertions: assert.Assertions,
		name:       name,
		str:        property.String(),
	}
}

// String looks up an element in a slice expecting it to be a string.
func (assert SliceAssertions) String(idx int) StringAssertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.String {
		assert.Fatal("element %d is not a string", idx)
	}

	return StringAssertions{
		Assertions: assert.Assertions,
		name:       name,
		str:        property.String(),
	}
}

// Blank generates an error if the string is not empty with something other than whitespace.
func (assert StringAssertions) Blank() {
	assert.t.Helper()

	if strings.TrimSpace(assert.str) != "" {
		assert.Failed(`%s is not blank; is "%s"`, assert.name, assert.str)
	}
}

// NotBlank generates an error if the string is empty or contains only whitespace.
func (assert StringAssertions) NotBlank() {
	assert.t.Helper()

	if strings.TrimSpace(assert.str) == "" {
		assert.Failed(`%s is blank`, assert.name)
	}
}

// IsBlank is an alias for Blank.
func (assert StringAssertions) IsBlank() {
	assert.t.Helper()

	if strings.TrimSpace(assert.str) != "" {
		assert.Failed(`%s is not blank; is "%s"`, assert.name, assert.str)
	}
}

// IsNotBlank is an alias for NotBlank.
func (assert StringAssertions) IsNotBlank() {
	assert.t.Helper()

	if strings.TrimSpace(assert.str) == "" {
		assert.Failed(`%s is blank`, assert.name)
	}
}

// Empty generates an error if the string contains any characters, including whitespace.
func (assert StringAssertions) Empty() {
	assert.t.Helper()

	if assert.str != "" {
		assert.Failed(`%s is not empty; is "%s"`, assert.name, assert.str)
	}
}

// NotEmpty generates an error if the string does not contain any characters, including whitespace.
func (assert StringAssertions) NotEmpty() {
	assert.t.Helper()

	if assert.str == "" {
		assert.Failed(`%s is empty"`, assert.name)
	}
}

// Equals generates an error if the string value isn't the same as other.
func (assert StringAssertions) Equals(other string) {
	assert.t.Helper()

	if assert.str != other {
		assert.Failed(`expected %s to be "%s" but it was "%s"`, assert.name, other, assert.str)
	}
}

// NotEquals generates an error if the string value is the same as the other.
func (assert StringAssertions) NotEquals(other string) {
	assert.t.Helper()

	if assert.str == other {
		assert.Failed(`expected %s to not be "%s"`, assert.name, other)
	}
}

// Contains generates an error if the other value isn't present somewhere in the string.
func (assert StringAssertions) Contains(other string) {
	assert.t.Helper()

	if !strings.Contains(assert.str, other) {
		assert.Failed(`expected %s to contain "%s" (is "%s")`, assert.name, other, assert.str)
	}
}

// NotContains generates an error if the other value is present somewhere in the string.
func (assert StringAssertions) NotContains(other string) {
	assert.t.Helper()

	if strings.Contains(assert.str, other) {
		assert.Failed(`expected %s to not contain "%s" (is "%s")`, assert.name, other, assert.str)
	}
}

// EndsWith generates an error if the other value isn't present at the end of the string.
func (assert StringAssertions) EndsWith(other string) {
	assert.t.Helper()

	if !strings.HasSuffix(assert.str, other) {
		assert.Failed(`expected %s to have "%s" at the end (is "%s")`, assert.name, other, assert.str)
	}
}

// StartsWith generates an error if the other value isn't present at the beginning of the string.
func (assert StringAssertions) StartsWith(other string) {
	assert.t.Helper()

	if strings.HasPrefix(assert.str, other) {
		assert.Failed(`expected %s to have "%s" at the beginning (is "%s")`, assert.name, other, assert.str)
	}
}

// Matches generates an error if the string value doesn't match the regular expression.
func (assert StringAssertions) Matches(expr string) {
	assert.t.Helper()

	matched, err := regexp.MatchString(expr, assert.str)
	if err != nil {
		assert.Fatal(`invalid regular expression "%s"`, expr)
	}

	if !matched {
		assert.Failed(`expected %s to match "%s" (is %s)`, assert.name, expr, assert.str)
	}
}

// NotMatches generates an error if the string value matches the regular expression.
func (assert StringAssertions) NotMatches(expr string) {
	assert.t.Helper()

	matched, err := regexp.MatchString(expr, assert.str)
	if err != nil {
		assert.Fatal(`invalid regular expression "%s"`, expr)
	}

	if matched {
		assert.Failed(`expected %s not to match "%s" (is %s)`, assert.name, expr, assert.str)
	}
}
