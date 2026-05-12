package doublecheck

import "sync"

type Singleton struct {
	name string
}

var locker sync.Mutex

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		locker.Lock()
		if instance == nil {
			instance = &Singleton{"double-check"}
		}
		locker.Unlock()
	}
	return instance
}
