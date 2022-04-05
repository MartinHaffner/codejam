package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func processTestCase(testCase int, lines []string) {
	fmt.Printf("Case #%d: \n", testCase)
	for _, line := range lines {
		println(line)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// max. line size: 100.000 bags * 5 chars (4 digits + space)
	const maxCapacity = 100000 * 5
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	first_line := scanner.Text()
	testCases, _ := strconv.Atoi(first_line)
	linesPerCase := 2 /* todo adapt */
	for testCase := 1; testCase <= testCases; testCase++ {
		lines := make([]string, linesPerCase)
		for line := 0; line < linesPerCase; line++ {
			scanner.Scan()
			lines[line] = scanner.Text()
		}
		processTestCase(testCase, lines)
	}
}
