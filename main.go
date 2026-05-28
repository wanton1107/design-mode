package main

import (
	"design-mode/decorator"
	"fmt"
)

func main() {
	drink := decorator.Factor(decorator.NewCoffee(10))
	drink = decorator.AddMilkDecorator(drink)
	drink = decorator.AddSugarDecorator(drink)
	fmt.Println("Coffee cost:", drink.Cost())
}
