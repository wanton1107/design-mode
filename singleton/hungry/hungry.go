package hungry

type Singleton struct {
	name string
}

var instance *Singleton

func init() {
	instance = &Singleton{"hungry"}
}

func GetInstance() *Singleton {
	return instance
}
