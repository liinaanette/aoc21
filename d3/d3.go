package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("task1 answer: ", task1())
	fmt.Println("task2 answer: ", task2())
}

func task1() int {
	file := openFile()
	transposed := make([]string, 0)

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for i, bit := range line {
			if len(transposed) <= i {
				transposed = append(transposed, "")
			}
			transposed[i] += bit
		}
	}

	gamma := ""
	epsilon := ""
	for _, bitRow := range transposed {
		oneLen := strings.Count(bitRow, "1")

		bit := 0
		if oneLen > len(bitRow)/2 {
			bit = 1
		}
		gamma += strconv.Itoa(bit)
		epsilon += strconv.Itoa(1 - bit)

	}

	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	check(err)
	epislonInt, err := strconv.ParseInt(epsilon, 2, 64)
	check(err)

	return int(gammaInt) * int(epislonInt)
}

func task2() int {
	return 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile() *os.File {
	f, err := os.Open("input.txt")
	check(err)
	return f
}
