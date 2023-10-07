package main

import (
	"fmt"
)

type IPizza interface {
	getPrice() int
}
type VeggieMania struct {
	price int
}

func (v *VeggieMania) getPrice() int {
	return v.price
}

type TomatoTopping struct {
	pizza IPizza
	price int
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + t.price
}

type CheeseTopping struct {
	pizza IPizza
	price int
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + c.price
}
func main() {
	veggieMania := &VeggieMania{price: 10}
	tomatoTopping := &TomatoTopping{pizza: veggieMania, price: 2}
	cheeseTopping := &CheeseTopping{pizza: tomatoTopping, price: 3}
	totalPrice := cheeseTopping.getPrice()
	fmt.Println("Total price:", totalPrice)
}
