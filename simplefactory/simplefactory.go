package simplefactory

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

func CreateAnimal(animalType string) Animal {
	switch animalType {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}
