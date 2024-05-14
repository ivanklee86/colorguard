package colorguard

import (
	"bytes"
	"testing"
)

func TestColorguardHappyPath(t *testing.T) {
	b := bytes.NewBufferString("")

	colorguard := New()
	colorguard.Out = b
	colorguard.Err = b
}
