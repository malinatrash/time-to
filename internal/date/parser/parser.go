package parser

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/malinatrash/time-to/pkg/checker"
)

func ParseFlag() (measureFlag *string, dateFlag *string) {
	measure := flag.String("m", "d", "unit of measure to use")
	flagDate := flag.String("d", "", "date until count")
	flag.Parse()

	if *flagDate == "" {
		fmt.Println("date is required")
		return
	}

	return measure, flagDate
}

func ParseDate(flagDate *string) time.Time {
	date := strings.Split(*flagDate, ".")

	day, err := strconv.Atoi(date[0])
	checker.Check(err)
	month, err := strconv.Atoi(date[1])
	checker.Check(err)

	year, err := strconv.Atoi(date[2])
	checker.Check(err)

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}
