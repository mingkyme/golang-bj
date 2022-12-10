package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var input string
var inputIndex int = 0
var inputSplit []string

func main() {
	input = `28
20
5`
	inputSplit = strings.Split(input, "\n")
	defer writer.Flush()

	var S, A, B int

	Readline(reader, &S)
	Readline(reader, &A)
	Readline(reader, &B)

	var needHeight = S - A
	if needHeight <= 0 {
		fmt.Fprintln(writer, 250)
		return
	}
	needCount := int(math.Ceil(float64(needHeight) / float64(B)))
	fmt.Fprintln(writer, 250+needCount*100)
}
func Readline(r io.Reader, a ...any) {
	if runtime.GOOS == "darwin" {
		buf := bytes.NewBufferString(inputSplit[inputIndex])
		inputIndex++
		fmt.Fscanln(buf, a...)
	} else {
		fmt.Fscanln(reader, a...)
	}
}
