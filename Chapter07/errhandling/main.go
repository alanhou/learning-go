package main

import (
	"bufio"
	"fmt"
	"io"
	"learninggo/Chapter07/errhandling/fib"
	"strings"
)

type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: 解决 p 过小导致的问题
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()

	printFileContents(f)
}
