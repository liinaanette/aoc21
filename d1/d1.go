package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("task1 answer: ", task1("input1.txt"))
	fmt.Println("task2 answer: ", task2("input1.txt"))
}

func task1(fileName string) int {
	file := openFile(fileName)

	bigger := 0
	last := 0
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		check(err)

		if last != 0 && current > last {
			bigger++
		}

		last = current
	}
	return bigger

}

func task2(fileName string) int {
	file := openFile(fileName)

	bigger := 0
	lastSum := 0
	lastThree := make([]int, 0)
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		check(err)

		lastThree = append(lastThree, current)

		if len(lastThree) == 3 {
			currentSum := sum(lastThree)
			if currentSum > lastSum && lastSum != 0 {
				bigger++
			}
			lastSum = currentSum
			lastThree = lastThree[1:]
		}
	}
	return bigger

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	check(err)
	return f
}
