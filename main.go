package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	minDelay = flag.Duration(
		"min-delay", 5*time.Millisecond,
		"minimum delay between printing a character")
	maxDelay = flag.Duration(
		"max-delay",
		200*time.Millisecond,
		"maximum delay between printing a character")
)

func TextReader() io.Reader {
	if flag.NArg() > 0 {
		return strings.NewReader(flag.Arg(0))
	}
	return os.Stdin
}

func DelayedPrint(r rune) {
	d := *minDelay + time.Duration(rand.Int63n(int64(*maxDelay-*minDelay)))
	time.Sleep(d)
	print(string(r))
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [text]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Print provided [text] with random delays between characters.")
		fmt.Fprintf(os.Stderr, " If [text] is not provided, then prints stdin.\n\n")
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	r := bufio.NewReader(TextReader())
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading input: %+v", err)
			os.Exit(1)
		}
		DelayedPrint(c)
	}
}
