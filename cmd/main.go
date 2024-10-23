package main

import (
	"github.com/malinatrash/time-to/internal/date/calculator"
	"github.com/malinatrash/time-to/internal/date/parser"
)

func main() {
	measure, date := parser.ParseFlag()
	parsedDate := parser.ParseDate(date)
	calculator.CalculateDate(parsedDate, measure)
}
