// Package tort implements an simple assertions testing framework for Go.
package tort

import "testing"

// For initializes the Tort assertions for a test case or benchmark.
func For(tb testing.TB) Assertions {
	return NewAssertions(tb)
}
