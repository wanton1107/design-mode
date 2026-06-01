package main

import (
	"design-mode/flyweight"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		activity := flyweight.QueryActivity(1)
		fmt.Println(activity)
		time.Sleep(1 * time.Second)
	}
}	
 
