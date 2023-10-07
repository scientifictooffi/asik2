package main

import (
	"fmt"
)

type IProduct interface {
	getName() string
}

type Car struct {
	name string
}

func (c *Car) getName() string {
	return c.name
}

type Phone struct {
	name string
}

func (p *Phone) getName() string {
	return p.name
}

type IFactory interface {
	createProduct() IProduct
}

type CarFactory struct{}

func (c *CarFactory) createProduct() IProduct {
	return &Car{name: "Car"}
}

type PhoneFactory struct{}

func (p *PhoneFactory) createProduct() IProduct {
	return &Phone{name: "Phone"}
}

func main() {
	carFactory := &CarFactory{}

	car := carFactory.createProduct()

	fmt.Println(car.getName())

	phoneFactory := &PhoneFactory{}

	phone := phoneFactory.createProduct()

	fmt.Println(phone.getName())
}
