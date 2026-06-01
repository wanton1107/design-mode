package flyweight

import (
	"math/rand"
	"sync"
	"time"
)

var stock int32 = 0
var lock sync.Mutex

func GetStock() int32 {
	lock.Lock()
	defer lock.Unlock()
	return stock
}

func AddAndGet(count int32) int32 {
	lock.Lock()
	defer lock.Unlock()
	stock = stock + count
	return stock
}

func InitRedis() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				AddAndGet(rand.Int31n(10))
			}
		}
	}()
}

func init() {
	InitRedis()
}
