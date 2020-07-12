package tort


// ErrorAssertions test errors.
type ErrorAssertions struct {
	Assertions
	err error
}

// Error assists in validating errors occurred.
func (assert Assertions) Error(err error) ErrorAssertions {
	return ErrorAssertions{
		Assertions: assert,
		err: err,
	}
}

// IsNil generates an error message if the error isn't nil.
func (assert ErrorAssertions) IsNil() {
	if assert.err != nil {
		assert.Failed(`unexpected error "%s"`, assert.err)
	}
}

// IsNotNil generates an error message when the error is nil.
func (assert ErrorAssertions) IsNotNil() {
	if assert.err == nil {
		assert.Failed("expected error wasn't present")
	}
}

// Equal checks that the expected error was generated.
func (assert ErrorAssertions) Equal(expected error) {
	if assert.err != expected {
		assert.Failed(`expected error "%s", but got "%s"`, expected, assert.err)
	}
}

