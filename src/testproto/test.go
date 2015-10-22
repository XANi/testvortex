package testproto

import (
	"time"
)

// Single test structure, heavily based off https://testanything.org/tap-version-13-specification.html with some JUnit additions

type Test struct {
	Success bool `json:"success"`
	Description string `json:"description"`
	// test output, that should contain any message from failing test
	Output string `json:"out"`
	// specify what a given test suite expected, if supported by input
	Expected string `json:"expected"`
	// specify what a given test suite got, if supported by input
	Got string `json:"got"`
	// Duration, if given test suite can output it
	Duration time.Duration `json:"duration,omitempty"`
	// Test file or identifier
	TestFile string `json:"test_file,omitempty"`
	// Test line, if supported
	TestLine int `json:"test_line,omitempty"`
	// Skipped means "skipped because of reasons" like missing systems stuff or DB to run test
	// Should be not counted into failure count
	Skip bool `json:"skip,omitempty"`
	// Reason for skipping, if any
	SkipReason string `json:"skip,omitempty"`
	// Todo means test that are expected to fail (tested code not ready
	// Crashed means "no result, test suite exited" should be considered a failure
	Todo bool `json:"todo,omitempty"`
	TodoReason string `json:"todo_reason,omitempty"`
	Crash bool `json:"crash,omitempty"`
	// Extra data from test if applicable
	Details interface {} `json:"details,omitempty"`
}
