package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var input string
var inputIndex int = 0
var inputSplit []string

func main() {
	input = `2
7
I 16
I -5643
D -1
D 1
D 1
I 123
D -1
9
I -45
I 653
D 1
I -642
I 45
I 97
D 1
D -1
I 333`
	inputSplit = strings.Split(input, "\n")
	defer writer.Flush()

	var T int

	Readline(reader, &T)
	for i := 0; i < T; i++ {
		var Q int
		Readline(reader, &Q)
		var array []int
		for j := 0; j < Q; j++ {
			var command string
			var num int
			Readline(reader, &command, &num)
			if command == "I" {
				array = append(array, num)
				sort.Slice(array, func(i, j int) bool {
					return array[i] < array[j]
				})
			} else if command == "D" {
				if len(array) == 0 {
					continue
				}
				if num == 1 {
					array = array[:len(array)-1]
				} else if num == -1 {
					array = array[1:]
				}
			}
		}
		if len(array) == 0 {
			fmt.Fprintln(writer, "EMPTY")
		} else {
			fmt.Fprintln(writer, array[len(array)-1], array[0])
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
