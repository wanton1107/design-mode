package builder

type Builder interface {
	Reset() Builder
	SetName(name string) Builder
	SetPrice(price float64) Builder
	SetTea(tea string) Builder
	SetSugar(sugar string) Builder
	SetIce(ice string) Builder
	Build() MilkTea
}

type MilkTeaBuilder struct {
	milkTea MilkTea
}

func (b *MilkTeaBuilder) Reset() Builder {
	b.milkTea = MilkTea{}
	return b
}

func (b *MilkTeaBuilder) SetName(name string) Builder {
	b.milkTea.Name = name
	return b
}

func (b *MilkTeaBuilder) SetPrice(price float64) Builder {
	b.milkTea.Price = price
	return b
}

func (b *MilkTeaBuilder) SetTea(tea string) Builder {
	b.milkTea.Tea = tea
	return b
}

func (b *MilkTeaBuilder) SetSugar(sugar string) Builder {
	b.milkTea.Sugar = sugar
	return b
}

func (b *MilkTeaBuilder) SetIce(ice string) Builder {
	b.milkTea.Ice = ice
	return b
}

func (b *MilkTeaBuilder) Build() MilkTea {
	return b.milkTea
}
