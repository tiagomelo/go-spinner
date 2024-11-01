// Copyright (c) 2024 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package spinner

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

// spinner represents a spinner component.
type spinner struct {
	charset       []string      // charset is the set of characters used to render the spinner animation.
	title         string        // title is the text displayed next to the spinner animation.
	concludedChar string        // concludedChar is the character displayed when the spinner stops.
	frameRate     time.Duration // frameRate is the speed at which the spinner animation is rendered.
	runChan       chan struct{} // runChan is a channel used to control the spinner animation goroutine.
	startOnce     sync.Once     // startOnce ensures the spinner animation goroutine is started only once.
	stopOnce      sync.Once     // stopOnce ensures the spinner animation goroutine is stopped only once.
	writer        io.Writer     // writer is the writer used to render the spinner animation.
	isTerminal    bool          // isTerminal is true if the writer is a terminal.
}

// patchSpinner applies the provided options to a spinner.
func patchSpinner(spinner *spinner, options []Option) *spinner {
	for _, option := range options {
		option(spinner)
	}
	spinner.isTerminal = isTerminal(spinner.writer)
	return spinner
}

// isTerminal returns true if the provided writer is a terminal.
func isTerminal(w io.Writer) bool {
	if file, ok := w.(*os.File); ok {
		fd := file.Fd()
		return term.IsTerminal(int(fd))
	}
	return false
}

// New creates a new spinner component with the provided options.
// If no options are provided, the spinner will use the default options:
//   - charset: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
//   - concludedChar: "✔"
//   - speed: 150ms
//   - writer: os.Stdout
func New(title string, options ...Option) *spinner {
	const defaultFrameRate = 150 * time.Millisecond
	writer := os.Stdout
	spinner := &spinner{
		title:         title,
		charset:       []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		concludedChar: "✔",
		frameRate:     defaultFrameRate,
		runChan:       make(chan struct{}),
		writer:        writer,
		isTerminal:    isTerminal(writer),
	}
	return patchSpinner(spinner, options)
}

// Start starts the spinner animation.
// The spinner will only start if the writer is a terminal.
func (s *spinner) Start() {
	if !s.isTerminal {
		return
	}
	s.startOnce.Do(func() {
		go s.writeAnimation()
	})
}

// writeAnimation writes the spinner animation to the terminal.
func (s *spinner) writeAnimation() {
	s.animate()
	for {
		select {
		case <-s.runChan:
			return
		default:
			s.animate()
		}
	}
}

// animate writes the spinner animation to the terminal.
func (s *spinner) animate() {
	var out string
	for i := 0; i < len(s.charset); i++ {
		out = fmt.Sprintf("%s %s", s.charset[i], s.title)
		fmt.Print(out)
		time.Sleep(s.frameRate)
		s.clearLine()
	}
}

// clearLine clears the current line in the terminal.
func (s *spinner) clearLine() {
	if s.isTerminal {
		// Constants for ANSI escape sequences
		const (
			clearLine    = "\033[2K" // clears the current line in the terminal.
			moveCursorUp = "\033[1A" // moves the cursor up one line.
		)
		fmt.Fprint(s.writer, clearLine)
		fmt.Fprintln(s.writer)
		fmt.Fprint(s.writer, moveCursorUp)
	}
}

// Stop stops the spinner animation.
func (s *spinner) Stop() {
	s.stopOnce.Do(func() {
		close(s.runChan)
		s.clearLine()
		fmt.Fprintf(s.writer, "%s %s\n", s.concludedChar, s.title)
	})
}
