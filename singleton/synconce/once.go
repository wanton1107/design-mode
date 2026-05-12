package synconce

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var once sync.Once

var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		fmt.Println("创建一次")
	})
	return instance
}
