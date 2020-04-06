package main

import (
	"fmt"
)

func main() {
	var b beverage
	b = newDarkRoast()
	b = newMocha(b)
	b = newMocha(b)
	b = newWhip(b)
	b.describe()
	fmt.Printf("cost:%+v\n", b.cost())
}
