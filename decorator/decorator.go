package decorator

type Factor interface {
	Cost() int32
}

type Coffee struct {
	price int32
}

func NewCoffee(price int32) *Coffee {
	return &Coffee{price: price}
}

func (c *Coffee) Cost() int32 {
	return c.price
}

type MilkDecorator struct {
	factor Factor
}

func AddMilkDecorator(factor Factor) Factor {
	return &MilkDecorator{factor: factor}
}

func (m *MilkDecorator) Cost() int32 {
	return m.factor.Cost() + 10
}

type SugarDecorator struct {
	factor Factor
}

func AddSugarDecorator(factor Factor) Factor {
	return &SugarDecorator{factor: factor}
}


func (s *SugarDecorator) Cost() int32 {
	return s.factor.Cost() + 5
}
