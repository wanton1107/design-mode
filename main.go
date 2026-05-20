package main

import (
	"design-mode/prototype"
	
)

func main() {
	gasCar := &prototype.GasCar{
		Brand:  "宝马",
		Fuel:   "汽油",
		Color:  "红色",
		Engine: "1.5T",
	}
	electricCar := &prototype.ElectricCar{
		Brand:  "特斯拉",
		Fuel:   "电力",
		Color:  "蓝色",
		Engine: "1.5T",
	}
	for i := 0; i < 10; i++ {
		gasCarCopy := gasCar.Clone()
		electricCarCopy := electricCar.Clone()
		fmt.Println(gasCarCopy)
		fmt.Println(electricCarCopy)
	}
}
