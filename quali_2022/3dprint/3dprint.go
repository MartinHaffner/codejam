package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = [...]string{"cyan", "magenta", "yellow", "black"}

const fullValue = 1000000

type PrinterFillState map[string]int

func getMinInkValues(fillStates []PrinterFillState) PrinterFillState {
	result := make(PrinterFillState)
	for _, color := range colors {
		minValue := fullValue
		for _, fillState := range fillStates {
			currentValue := fillState[color]
			if currentValue < minValue {
				minValue = currentValue
			}
		}
		result[color] = minValue
	}
	return result
}

func getTotalInkValue(fillState PrinterFillState) int {
	result := 0
	for _, color := range colors {
		result += fillState[color]
	}
	return result
}

func adaptSolution(fillState PrinterFillState) PrinterFillState {
	result := make(PrinterFillState)
	for color, fill := range fillState {
		result[color] = fill
	}
	currentSum := getTotalInkValue(result)
	for _, color := range colors {
		overhead := currentSum - fullValue
		if overhead > 0 {
			if overhead > fillState[color] {
				result[color] = 0
			} else {
				result[color] = fillState[color] - overhead
			}
		}
		currentSum = getTotalInkValue(result)
	}
	return result
}

func findSolution(fillStates []PrinterFillState) string {
	result := ""
	minInkValues := getMinInkValues(fillStates)
	valueSum := getTotalInkValue(minInkValues)
	if valueSum < fullValue {
		return "IMPOSSIBLE"
	}
	adapted := adaptSolution(minInkValues)
	for _, color := range colors {
		result += fmt.Sprintf("%d ", adapted[color])
	}
	return strings.TrimSpace(result)
}

func processTestCase(testCase int, lines []string) {
	fillStates := make([]PrinterFillState, 0)
	for _, line := range lines {
		lineValues := strings.Split(line, " ")
		fillState := make(PrinterFillState)
		for colNum, color := range colors {
			colValue, _ := strconv.Atoi(lineValues[colNum])
			fillState[color] = colValue
		}
		fillStates = append(fillStates, fillState)
	}
	fmt.Printf("Case #%d: %s\n", testCase, findSolution(fillStates))
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 100000 * 5
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	first_line := scanner.Text()
	testCases, _ := strconv.Atoi(first_line)
	linesPerCase := 3
	for testCase := 1; testCase <= testCases; testCase++ {
		lines := make([]string, linesPerCase)
		for line := 0; line < linesPerCase; line++ {
			scanner.Scan()
			lines[line] = scanner.Text()
		}
		processTestCase(testCase, lines)
	}
}
