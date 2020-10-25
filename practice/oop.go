package practice

import "fmt"

type Item struct {
	Name string
	Price float64
	Quantity int
}

type quantity int
type dozen []quantity

func (i Item) Cost() float64 {
	return i.Price * float64(i.Quantity)
}

func CustomizedType() {
	q := quantity(1)
	d := dozen{q}
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}

	fmt.Println(d)
}
