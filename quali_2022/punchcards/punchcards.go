package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFirstLine(c int, isTopLine bool) string {
	var result string
	if isTopLine {
		result = ".."
	} else {
		result = "+-"
	}
	for col := 2; col <= c; col++ {
		result += "+-"
	}
	result += "+"
	return result
}

func getSecondLine(c int, isTopLine bool) string {
	var result string
	if isTopLine {
		result = ".."
	} else {
		result = "|."
	}
	for col := 2; col <= c; col++ {
		result += "|."
	}
	result += "|"
	return result
}

func outputPunchCard(r int, c int) {
	for row := 1; row <= r; row++ {
		isTopLine := row == 1
		fmt.Println(getFirstLine(c, isTopLine))
		fmt.Println(getSecondLine(c, isTopLine))
	}
	fmt.Println(getFirstLine(c, false))
}

func processTestCase(testCase int, lines []string) {
	fmt.Printf("Case #%d:\n", testCase)
	params := strings.Split(lines[0], " ")
	r, _ := strconv.Atoi(params[0])
	c, _ := strconv.Atoi(params[1])
	outputPunchCard(r, c)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 100000 * 5
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	first_line := scanner.Text()
	testCases, _ := strconv.Atoi(first_line)
	linesPerCase := 1 /* todo adapt */
	for testCase := 1; testCase <= testCases; testCase++ {
		lines := make([]string, linesPerCase)
		for line := 0; line < linesPerCase; line++ {
			scanner.Scan()
			lines[line] = scanner.Text()
		}
		processTestCase(testCase, lines)
	}
}
