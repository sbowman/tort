package tort

import (
	"fmt"
	"strings"
	"testing"
)

// Assertions are the base set of assertions.
type Assertions struct {
	t   testing.TB
	msg string
}

// A provides an alias shortcut for assertions.  Makes it easier to use the With function in test
// cases.
type A struct {
	Assertions
}

// NewAssertions creates a new Assertions object to use for a test case.
func NewAssertions(t testing.TB) Assertions {
	return Assertions{
		t: t,
	}
}

// IsNil checks if the value is nil.  If not, generates an error.
func (assert Assertions) IsNil(val interface{}, msg ...string) {
	assert.t.Helper()

	if val != nil {
		if len(msg) == 0 {
			assert.Failed(`unexpected value "%s"`, val)
			return
		}

		assert.Failed(`%s: %s`, strings.Join(msg, " "), val)
	}
}

// IsNotNil checks that the value is not nil.  If it is, generates an error.
func (assert Assertions) IsNotNil(val interface{}, msg ...string) {
	assert.t.Helper()

	if val == nil {
		if len(msg) == 0 {
			assert.Failed(`unexpected value "%s" was not present`, val)
			return
		}

		assert.Failed(`%s: %s`, strings.Join(msg, " "), val)
	}
}

// IsTrue confirms the value is true.  If not, generates an error.
func (assert Assertions) IsTrue(val bool, msg ...string) {
	assert.t.Helper()

	if !val {
		if len(msg) == 0 {
			assert.Failed("value was false; should have been true")
			return
		}

		assert.Failed(`%s: %s`, strings.Join(msg, " "), val)
	}
}

// IsFalse confirms the value is false.  If not, generates an error.
func (assert Assertions) IsFalse(val bool, msg ...string) {
	assert.t.Helper()

	if val {
		if len(msg) == 0 {
			assert.Failed("value was true; should have been false")
			return
		}

		assert.Failed(`%s: %s`, strings.Join(msg, " "), val)
	}
}

// When describes something about the assertion, e.g. assert.When("creating a user").
func (assert Assertions) When(msg ...string) Assertions {
	if len(msg) == 1 {
		assert.msg = msg[0]
	} else {
		assert.msg = strings.Join(msg, " ")
	}
	return assert
}

// With is like When, except that the assertion is passed to a function for processing.
func (assert Assertions) With(msg string, fn func(A)) {
	assert.msg = msg
	fn(A{assert})
}

// Failed outputs the final error message to *testing.T.
func (assert Assertions) Failed(format string, args ...interface{}) {
	assert.t.Helper()

	local := fmt.Sprintf(format, args...)

	if assert.msg != "" {
		assert.t.Errorf("when %s, %s", assert.msg, local)
		return
	}

	assert.t.Error(local)
}

// Fatal outputs the final error message to *testing.T and exits.
func (assert Assertions) Fatal(format string, args ...interface{}) {
	assert.t.Helper()

	local := fmt.Sprintf(format, args...)

	if assert.msg != "" {
		assert.t.Fatalf("%s when %s", local, assert.msg)
		return
	}

	assert.t.Fatal(local)
}
