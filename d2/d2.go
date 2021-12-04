package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("task1 answer: ", task1("input.txt"))
	fmt.Println("task2 answer: ", task2("input.txt"))
}

func task1(fileName string) int {
	file := openFile(fileName)
	forward := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		currentNumber, err := strconv.Atoi(line[1])
		check(err)

		switch currentOrder := line[0]; currentOrder {
		case "forward":
			forward += currentNumber
		case "up":
			depth -= currentNumber
		case "down":
			depth += currentNumber
		}
	}

	return forward * depth
}

func task2(fileName string) int {
	file := openFile(fileName)
	forward := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		currentNumber, err := strconv.Atoi(line[1])
		check(err)

		switch currentOrder := line[0]; currentOrder {
		case "forward":
			forward += currentNumber
			depth += aim * currentNumber
		case "up":
			aim -= currentNumber
		case "down":
			aim += currentNumber
		}
	}

	return forward * depth
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	check(err)
	return f
}
