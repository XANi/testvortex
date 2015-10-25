package testproto

import (
	"io"
	"time"
)

// test suite

type TestSuite struct {
	// testsuite name
	Name string `json:"name"`
	// description, if supported
	Description string `json:"description"`
	// Duration, if given test suite can output it
	Duration time.Duration `json:"duration,omitempty"`
	// tests run, in order
	Tests []Test `json:"test"`
	// Does whole test suite suceeded
	Success bool `json:"success"`
	Skip    bool `json:"skip,omitempty"`
	// Reason for skipping, if any
	SkipReason string `json:"skip,omitempty"`
	// Todo means test that are expected to fail (tested code not ready
	// Crashed means "no result, test suite exited" should be considered a failure
	Todo       bool   `json:"todo,omitempty"`
	TodoReason string `json:"todo_reason,omitempty"`
	Crash      bool   `json:"crash,omitempty"`
	// Extra data from test if applicable
	Details interface{} `json:"details,omitempty"`
}

type TestSuiteProto interface {
	Load(io.Reader) (err error)
	Save(io.Writer) (err error)
}
