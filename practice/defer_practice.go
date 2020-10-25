package practice

import "fmt"

func F1() {
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Println("f2 - difer")
}
