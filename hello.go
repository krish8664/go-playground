package main

import (
	"fmt"
)

func pointers(iptr *int) {
	fmt.Println("Address: ", iptr)
	fmt.Println("Value: ", *iptr)
	return
}

func playingWithPointerStructs() {
	r := rectangle{width: 10, height: 20}
	fmt.Println(r)

	r.canMutate(5, 5)
	fmt.Println(r)

	r.cantMutate(6, 6)
	fmt.Println(r)

	//works same as above
	r1 := &rectangle{width: 15, height: 25}
	fmt.Println(r1)

	r1.canMutate(5, 5)
	fmt.Println(r1)

	r1.cantMutate(6, 6)
	fmt.Println(r1)

}

func measure(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func returnErrorOn10(i int) (int, error) {
	if i == 10 {
		return 0, &myErrors{10, "Screwed"}
	}

	return 10, nil
}

func main() {
	fmt.Println("Hello world")

	someInt := 10
	pointers(&someInt)

	playingWithPointerStructs()

	r := rectangle{width: 10, height: 10}
	c := circle{radius: 1}
	measure(r)
	measure(c)

	if _, err := returnErrorOn10(10); err != nil {
		//using type assertion
		a, b := err.(*myErrors)
		fmt.Println(a)
		fmt.Println(b)
	}

}
