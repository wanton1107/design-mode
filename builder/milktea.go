package builder

import "fmt"

type MilkTea struct {
	Name  string
	Price float64
	Tea   string
	Sugar string
	Ice   string
}

func (m *MilkTea) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Tea: %s, Sugar: %s, Ice: %s\n", m.Name, m.Price, m.Tea, m.Sugar, m.Ice)
}
