package tort


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
		err: err,
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

		assert.Failed(`%s: %s`, msg[0], assert.err)
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

		assert.Failed(msg[0])
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
