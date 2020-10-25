package practice

import "fmt"

func SwitchCondition() {
	c := "a"
	switch c {
	case "a":
		fmt.Println("c is a")
	case "b":
		fmt.Println("c is b")
	default:
		fmt.Println("c is nothing")
	}
}
