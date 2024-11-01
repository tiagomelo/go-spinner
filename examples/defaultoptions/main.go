// Copyright (c) 2024 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/tiagomelo/go-spinner"
)

func main() {
	sp := spinner.New("executing task 1...")
	sp.Start()
	time.Sleep(3 * time.Second)
	sp.Stop()
	sp = spinner.New("executing task 2...")
	sp.Start()
	time.Sleep(3 * time.Second)
	sp.Stop()
	fmt.Println("done")
}
