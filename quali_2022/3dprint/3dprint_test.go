package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinInkValues(t *testing.T) {
	printers := []PrinterFillState{
		PrinterFillState{"cyan": 300000, "magenta": 200000, "yellow": 200000, "black": 300000},
		PrinterFillState{"cyan": 300000, "magenta": 200000, "yellow": 500000, "black": 500000},
		PrinterFillState{"cyan": 300000, "magenta": 500000, "yellow": 300000, "black": 200000}}
	expected := PrinterFillState{"cyan": 300000, "magenta": 200000, "yellow": 200000, "black": 200000}
	actual := getMinInkValues(printers)
	assert.Equal(t, expected, actual)
}

func TestAdaptSolution(t *testing.T) {
	actual := adaptSolution(PrinterFillState{"cyan": 68763,
		"magenta": 148041,
		"yellow": 178147,
		"black": 714381})
	expected := PrinterFillState{"cyan": 0,
		"magenta": 107472,
		"yellow": 178147,
		"black": 714381}
	assert.Equal(t, expected, actual)
}
