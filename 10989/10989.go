package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
	input = `10
5
2
3
1
4
2
3
5
1
7`
	inputSplit = strings.Split(input, "\n")
	defer writer.Flush()

	count := 0
	Readline(reader, &count)

	array := make([]int, 10001)
	var num int
	for i := 0; i < count; i++ {
		Readline(reader, &num)

		array[num]++
	}

	for i := 0; i < 10001; i++ {
		for j := 0; j < array[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
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
