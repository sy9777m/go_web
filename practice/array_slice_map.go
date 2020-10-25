package practice

import "fmt"

func Declaration() {
	a := [5]int {1, 2, 3, 4, 5}
	b := []int {1, 2, 3, 4, 5, 6}
	 c := map[string]int {
		"one": 1,
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
