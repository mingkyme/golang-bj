package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	count := 0
	fmt.Fscanf(reader, "%d\n", &count)

	array := make([]int, 10001)
	var num int
	for i := 0; i < count; i++ {
		fmt.Fscanf(reader, "%d\n", &num)

		array[num]++
	}

	for i := 0; i < 10001; i++ {
		for j := 0; j < array[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}
