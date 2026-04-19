package ui

// Option represents a single selectable item in a prompt.
type Option struct {
	Label string
	Value string
}

// SelectResult is the outcome of a Select call.
type SelectResult struct {
	Index int
	Value string
	Label string
}