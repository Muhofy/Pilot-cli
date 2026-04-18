package ui

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Spinner is an animated terminal spinner.
type Spinner struct {
	msg    string
	frames []string
	stop   chan struct{}
	wg     sync.WaitGroup
}

var spinFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// NewSpinner creates a new spinner with the given message.
func NewSpinner(msg string) *Spinner {
	return &Spinner{
		msg:    msg,
		frames: spinFrames,
		stop:   make(chan struct{}),
	}
}

// Start begins the spinner animation in a goroutine.
func (s *Spinner) Start() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		i := 0
		for {
			select {
			case <-s.stop:
				fmt.Print("\r\033[2K") // clear line
				return
			default:
				color.New(color.FgCyan).Printf("\r  %s  %s", s.frames[i%len(s.frames)], s.msg)
				i++
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
}

// Stop halts the spinner and clears the line.
func (s *Spinner) Stop() {
	close(s.stop)
	s.wg.Wait()
}

// StopWithSuccess halts the spinner and prints a success message.
func (s *Spinner) StopWithSuccess(msg string) {
	s.Stop()
	Success(msg)
}

// StopWithError halts the spinner and prints an error message.
func (s *Spinner) StopWithError(msg string) {
	s.Stop()
	Error(msg)
}