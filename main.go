package main

import (
	"design-mode/singleton"
	"fmt"
	"time"
)

// TIP hello
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			instance := singleton.GetInstance()
			fmt.Printf("instance:%p\n", instance)
		}()
	}
	time.Sleep(1 * time.Second)
}
