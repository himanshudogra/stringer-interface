package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type miliLiters float64

func (m miliLiters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func main() {
	fmt.Println(Gallons(1.65))
	fmt.Println(Liters(4.78))
	fmt.Println(miliLiters(7.099))
}
