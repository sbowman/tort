package tort

import (
	"fmt"
	"strings"
	"testing"
)

// Assertions are the base set of assertions.
type Assertions struct {
	t testing.TB
	msg string
}

// NewAssertiona creates a new Assertions object to use for a test case.
func NewAssertions(t testing.TB) Assertions {
	return Assertions{
		t: t,
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
