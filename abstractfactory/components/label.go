package components

import "fmt"

type Label interface {
	Render()
}

type AndroidLabel struct {
}

func (l *AndroidLabel) Render() {
	fmt.Println("Material Design Label Render")
}

type IOSLabel struct {
}

func (l *IOSLabel) Render() {
	fmt.Println("IOS Label Render")
}
