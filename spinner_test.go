// Copyright (c) 2024 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package spinner

import (
	"bytes"
	"io"
	"sync"
	"testing"
)

func TestSpinnerStartAndStop(t *testing.T) {
	s := New("Processing", WithWriter(io.Discard))

	// Using WaitGroup to confirm animation goroutine has started.
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.Start()
	}()

	// Ensure Start executes once and doesn't block further calls.
	wg.Wait()
	s.Start() // Should not start another instance due to sync.Once

	// Stop the spinner and ensure it stops cleanly.
	s.Stop()
	s.Stop() // Should not stop twice due to sync.Once

	select {
	case <-s.runChan:
		// Expected behavior: channel should be closed after Stop.
	default:
		t.Errorf("Expected runChan to be closed, but it was open")
	}
}

func TestSpinnerOutput(t *testing.T) {
	// Create mock terminal writer to capture output
	mockWriter := &mockTerminalWriter{isTerminal: true}

	// Initialize spinner with custom writer
	s := New("Task", WithWriter(mockWriter))

	// Test clearLine output
	s.clearLine()

	var expectedClear string
	if mockWriter.String() != expectedClear {
		t.Errorf("Expected clearLine output %q, got %q", expectedClear, mockWriter.String())
	}
	// Reset buffer and test Stop output
	mockWriter.Reset()
	s.Stop()

	expectedStop := "✔ Task\n" // Include ANSI codes
	if mockWriter.String() != expectedStop {
		t.Errorf("Expected Stop output %q, got %q", expectedStop, mockWriter.String())
	}
}

type mockTerminalWriter struct {
	bytes.Buffer
	isTerminal bool
}

func (m *mockTerminalWriter) Fd() uintptr {
	return 0
}

func (m *mockTerminalWriter) IsTerminal() bool {
	return m.isTerminal
}
