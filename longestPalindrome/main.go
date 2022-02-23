package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("please input string ")
	input, _ := reader.ReadString('\n')

	input = "babad"

	Output := longestPalindrome(input)

	fmt.Printf("Input : %s", input)
	fmt.Println()
	fmt.Printf("Output : %s", Output)
	fmt.Scanln()
}

func longestPalindrome(s string) string {

	var longest int
	var result string

	stack := []byte{}
	queue := []byte{}

	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		queue = queueSlice(stack)

		fmt.Println(string(stack[:]))
		fmt.Println(string(queue[:]))

		if queue[0] != stack[i] {
			if len(stack)-1 > longest {
				longest = len(stack) - 1
				result = string(stack[:len(stack)-1])
			}
			stack = []byte{}
			queue = []byte{}
		}
	}

	return result
}

func queueSlice(stack []byte) []byte {
	queue := []byte{}

	for i := len(stack) - 1; i >= 0; i-- {
		queue = append(queue, stack[i])
	}
	return queue
}
