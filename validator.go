package tort

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidatorAssertions test validation errors from https://github.com/go-playground/validator.
type ValidatorAssertions struct {
	Assertions
	err validator.ValidationErrors
}

// Valid assists in validating validation errors,
func (assert Assertions) Valid(err error) ValidatorAssertions {
	assert.t.Helper()

	if ve, ok := err.(validator.ValidationErrors); ok {
		return ValidatorAssertions{
			Assertions: assert,
			err:        ve,
		}
	}

	assert.Failed("error is not a validator error")
	return ValidatorAssertions{}
}

// IsNil generates an error message if the error isn't nil.
func (assert ValidatorAssertions) IsNil(msg ...string) {
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
func (assert ValidatorAssertions) IsNotNil(msg ...string) {
	assert.t.Helper()

	if assert.err == nil {
		if len(msg) == 0 {
			assert.Failed("expected error wasn't present")
			return
		}

		assert.Failed(strings.Join(msg, " "))
	}
}

// For checks for a validation error matching the field and the kind of error, e.g. "required" or
// "min".  If kind is blank, confirms there was any error on the given field.
func (assert ValidatorAssertions) For(field, kind string) {
	assert.t.Helper()

	for _, fe := range assert.err {
		if fe.Field() == field {
			if kind == "" || fe.Tag() == kind {
				return
			}
		}
	}

	assert.Failed(`expected a validation error "%s" on field "%s"`, kind, field)
}

// NotFor checks that the validation error either doesn't exist, or isn't for the given field and
// type of error.  If kind is blank, just confirms there wasn't an error on the field.
func (assert ValidatorAssertions) NotFor(field, kind string) {
	assert.t.Helper()

	for _, fe := range assert.err {
		if fe.Field() == field {
			if kind == "" || fe.Tag() == kind {
				assert.Failed(`didn't expect a validation error "%s" on field "%s"`, kind, field)
			}
		}
	}
}

