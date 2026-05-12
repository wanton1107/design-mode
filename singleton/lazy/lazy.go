package lazy

import "sync"

type Singleton struct {
	name string
}

var locker sync.Mutex

var instance *Singleton

func GetInstance() *Singleton {
	locker.Lock()
	defer locker.Unlock()
	if instance == nil {
		instance = &Singleton{"lazy"}
	}
	return instance
}
