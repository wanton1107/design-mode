package components

import "fmt"

type Button interface {
	Render()
}

type AndroidButton struct {
}

func (b *AndroidButton) Render() {
	fmt.Println("Material Design Button Render")
}

type IOSButton struct {
}

func (b *IOSButton) Render() {
	fmt.Println("IOS Button Render")
}
