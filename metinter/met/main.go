package main

import "fmt"

type gallon float64
type kg int

func (k kg) double() int {
	return int(k * 2)
}

func (g gallon) quart() float64 {
	return float64(g * 4)
}

func main() {
	gal := gallon(5)
	fmt.Println(gal.quart())
	boy := kg(10)
	fmt.Println(boy.double())

}
