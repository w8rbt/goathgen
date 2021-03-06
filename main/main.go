/*
 * Copyright (c) 2016 Richard B Tilley <btilley@gatech.edu>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import "bufio"
import "fmt"
import "os"
import "time"

import "github.com/62726164/goathgen"

func main() {

	// Secret must be sent by stdin
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	now := time.Now()
	unixNow := now.Unix()

	var unixEpoch int64
	unixEpoch = 0

	var timeStep int64
	timeStep = 30

	if goathgen.Debug {
		fmt.Println("----- main -----")
		fmt.Printf("stdin %s\n", stdin.Text())
		fmt.Printf("unixNow: %d\n", unixNow)
		fmt.Printf("unixEpoch: %d\n", unixEpoch)
		fmt.Printf("timeStep: %d\n", timeStep)
		fmt.Println("----- main -----")
	}

	theTotp := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(stdin.Text(), unixNow, timeStep, unixEpoch)))
	fmt.Printf("%06d\n", theTotp)
}
