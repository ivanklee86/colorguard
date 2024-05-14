package colorguard

import (
	"io"
	"os"
)

type Config struct {
	NoExitCode bool
}

// Colorguard is the logic/orchestrator.
type Colorguard struct {
	*Config

	// Allow swapping out stdout/stderr for testing.
	Out io.Writer
	Err io.Writer
}

// New returns a new instance of Colorguard.
func New() *Colorguard {
	config := Config{}

	return &Colorguard{
		Config: &config,
		Out:    os.Stdout,
		Err:    os.Stdin,
	}
}
