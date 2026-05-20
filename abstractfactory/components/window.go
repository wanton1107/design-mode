package components

import "fmt"

type Window interface {
	Render()
}

type AndroidWindow struct {
}

func (w *AndroidWindow) Render() {
	fmt.Println("Material Design Window Render")
}

type IOSWindow struct {
}

func (w *IOSWindow) Render() {
	fmt.Println("IOS Window Render")
}
