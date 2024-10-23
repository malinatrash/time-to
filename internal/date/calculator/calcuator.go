package calculator

import (
	"fmt"
	"math"
	"time"
)

func CalculateDate(eventDate time.Time, measure *string) {
	now := time.Now()

	if !eventDate.After(now) {
		fmt.Println("date is in the past")
		return
	}

	switch *measure {
	case "y":
		fmt.Println(math.Ceil(eventDate.Sub(now).Hours() / 24 / 365))
	case "d":
		fmt.Println(math.Ceil(eventDate.Sub(now).Hours() / 24))
	case "h":
		fmt.Println(math.Ceil(eventDate.Sub(now).Hours()))
	case "m":
		fmt.Println(math.Ceil(eventDate.Sub(now).Minutes()))
	case "s":
		fmt.Println(eventDate.Sub(now).Seconds())
	}
}
