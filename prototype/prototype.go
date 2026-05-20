package prototype

type Prototype interface {
	Clone() Prototype
}

type GasCar struct {
	Brand  string
	Fuel   string
	Color  string
	Engine string
}

func (g *GasCar) Clone() Prototype {
	return &GasCar{Brand: g.Brand, Fuel: g.Fuel, Color: g.Color, Engine: g.Engine}
}

type ElectricCar struct {
	Brand  string
	Fuel   string
	Color  string
	Engine string
}

func (e *ElectricCar) Clone() Prototype {
	return &ElectricCar{Brand: e.Brand, Fuel: e.Fuel, Color: e.Color, Engine: e.Engine}
}
