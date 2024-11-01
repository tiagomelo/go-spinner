// Copyright (c) 2024 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tiagomelo/go-spinner"
)

func main() {
	sp := spinner.New("executing task 1...",
		spinner.WithClassicCharset(),
		spinner.WithConcludedChar("◆"),
		spinner.WithFrameRate(100*time.Millisecond),
	)
	sp.Start()
	time.Sleep(3 * time.Second)
	sp.Stop()

	sp = spinner.New("executing task 2...",
		spinner.WithCirclesCharset(),
		spinner.WithConcludedChar("★"),
		spinner.WithFrameRate(200*time.Millisecond),
	)
	sp.Start()
	time.Sleep(3 * time.Second)
	sp.Stop()

	// in this one we'll capture the output
	// to demonstrate the WithWriter option.
	// notice that in this case we're using a buffer,
	// so no animation will be displayed.
	var buf bytes.Buffer
	sp = spinner.New("executing task 3...",
		spinner.WithArrowsCharset(),
		spinner.WithConcludedChar("➜"),
		spinner.WithFrameRate(120*time.Millisecond),
		spinner.WithWriter(&buf),
	)
	sp.Start()
	time.Sleep(3 * time.Second)
	sp.Stop()
	fmt.Println("captured output for task 3:", buf.String())

	fmt.Println("done")
}
