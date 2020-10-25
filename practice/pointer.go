package practice

import "fmt"

func Pointer() {
	i := 1
	p := &i
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(p)
	fmt.Println(*p)
}

func Pointer1() {
	var p *int
	var pp **int

	i := 1
	p = &i
	pp = &p
	fmt.Println(i, p, pp)
	fmt.Println(i, *p, **pp)
}