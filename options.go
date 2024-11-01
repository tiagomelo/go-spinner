// Copyright (c) 2024 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package spinner

import (
	"io"
	"time"
)

// Option defines a function type for setting options on a spinner.
type Option func(*spinner)

// WithCharset sets the charset option for a spinner.
func WithCharset(charset ...string) Option {
	return func(s *spinner) {
		s.charset = charset
	}
}

// WithClassicCharset sets the spinner to use the classic charset.
func WithClassicCharset() Option {
	return WithCharset("|", "/", "-", "\\")
}

// WithArrowsCharset sets the spinner to use the arrows charset.
func WithArrowsCharset() Option {
	return WithCharset("←", "↖", "↑", "↗", "→", "↘", "↓", "↙")
}

// WithCirclesCharset sets the spinner to use the circles charset.
func WithCirclesCharset() Option {
	return WithCharset("◐", "◓", "◑", "◒")
}

// WithBlocksCharset sets the spinner to use the blocks charset.
func WithBlocksCharset() Option {
	return WithCharset("▖", "▘", "▝", "▗")
}

// WithConcludedChar sets the concludedChar option for a spinner.
func WithConcludedChar(concludedChar string) Option {
	return func(s *spinner) {
		s.concludedChar = concludedChar
	}
}

// WithFrameRate sets the frame rate option for a spinner.
func WithFrameRate(speed time.Duration) Option {
	return func(s *spinner) {
		s.frameRate = speed
	}
}

// WithWriter sets the writer option for a spinner.
func WithWriter(writer io.Writer) Option {
	return func(s *spinner) {
		s.writer = writer
	}
}
