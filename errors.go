package tort

import (
	"errors"
	"strings"
)

// ErrorAssertions test errors.
type ErrorAssertions struct {
	Assertions
	err error
}

// Error assists in validating errors occurred.
func (assert Assertions) Error(err error) ErrorAssertions {
	assert.t.Helper()

	return ErrorAssertions{
		Assertions: assert,
		err:        err,
	}
}

// IsNil generates an error message if the error isn't nil.
func (assert ErrorAssertions) IsNil(msg ...string) {
	assert.t.Helper()

	if assert.err != nil {
		if len(msg) == 0 {
			assert.Failed(`unexpected error "%s"`, assert.err)
			return
		}

		assert.Failed(`%s: %s`, strings.Join(msg, " "), assert.err)
	}
}

// IsNotNil generates an error message when the error is nil.
func (assert ErrorAssertions) IsNotNil(msg ...string) {
	assert.t.Helper()

	if assert.err == nil {
		if len(msg) == 0 {
			assert.Failed("expected error wasn't present")
			return
		}

		assert.Failed(strings.Join(msg, " "))
	}
}

// Equals checks that the expected error was generated.
func (assert ErrorAssertions) Equals(expected error) {
	assert.t.Helper()

	if assert.err != expected {
		assert.Failed(`expected error "%s" to be "%s"`, expected, assert.err)
	}
}

// NotEquals checks that the given error was not generated.
func (assert ErrorAssertions) NotEquals(expected error) {
	assert.t.Helper()

	if assert.err == expected {
		assert.Failed(`expected error "%s" not to be "%s"`, expected, assert.err)
	}
}

// Is checks the error is this kind of error.
func (assert ErrorAssertions) Is(expected error) {
	assert.t.Helper()

	if assert.err == nil || !errors.Is(assert.err, expected) {
		assert.Failed(`error "%s" is not "%s"`, assert.err, expected)
	}
}

// IsNot checks if the error is not this kind of error.
func (assert ErrorAssertions) IsNot(expected error) {
	assert.t.Helper()

	if errors.Is(assert.err, expected) {
		assert.Failed(`error "%s" is "%s"`, assert.err, expected)
	}
}
