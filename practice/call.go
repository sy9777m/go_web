package practice

import "fmt"

func calcVal(i int) {
	i += 1
}

func CallByValue() {
	i := 10
	calcVal(i)
	fmt.Println(i)

}

func calcRef(i *int)  {
	*i += 1
}

func CallByReference() {
	i := 10
	calcRef(&i)
	fmt.Println(i)
}
