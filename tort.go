// Package tort implements an simple assertions testing framework for Go.
package tort

import "testing"

// For initializes the Tort assertions for a test case.
func For(t *testing.T) Assertions {
	return NewAssertions(t)
}
