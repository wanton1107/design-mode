package methodfactory

type Animal interface {
	Shout() string
}

type Dog struct{}

func (d *Dog) Shout() string {
	return "汪汪汪"
}

type Cat struct{}

func (c *Cat) Shout() string {
	return "喵喵喵"
}

type Factory interface {
	CreateAnimal() Animal
}

type DogFactory struct{}

func (d *DogFactory) CreateAnimal() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (c *CatFactory) CreateAnimal() Animal {
	return &Cat{}
}
