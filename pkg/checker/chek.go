package checker

import "fmt"

func Check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
