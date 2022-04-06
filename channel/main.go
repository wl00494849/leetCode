package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Expresion interface {
	result() int
}

type StringMultiplication struct {
	expresion string
}

type StringAddition struct {
	expresion string
}

func main() {
	q := []string{"1+3", "0+2", "9+8", "9*2", "1*3"}
	res := solution(q)
	fmt.Printf("%v", res)
	fmt.Scanln()
}

func (sa StringAddition) result() int {
	num1, _ := strconv.Atoi(string(sa.expresion[0]))
	num2, _ := strconv.Atoi(string(sa.expresion[2]))
	return num1 + num2
}

func (sm StringMultiplication) result() int {
	num1, _ := strconv.Atoi(string(sm.expresion[0]))
	num2, _ := strconv.Atoi(string(sm.expresion[2]))
	return num1 * num2
}

func getResult(e Expresion) int {
	return e.result()
}

func generateVal(channel chan int, query string) {
	time.Sleep(500 * time.Millisecond)
	val := 0
	if query[1] == '+' {
		val = getResult(StringAddition{query})
	} else {
		val = getResult(StringMultiplication{query})
	}
	channel <- val
}

func solution(queries []string) []int {
	channel := make(chan int)
	for i := 0; i < 5; i++ {
		go generateVal(channel, queries[i])
		queries[i] = ""
	}

	sli := make([]int, 0)

	go func() {
		var i int
		for ch := range channel {
			sli = append(sli, ch)
			i++
			if i == 5 {
				sort.Ints(sli)
				close(channel)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	return sli
}
