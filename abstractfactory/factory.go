package abstractfactory

import "design-mode/abstractfactory/components"

type Factory interface {
	CreateButton() components.Button
	CreateLabel() components.Label
	CreateWindow() components.Window
}

type AndroidFactory struct {
}

func (f *AndroidFactory) CreateButton() components.Button {
	return &components.AndroidButton{}
}

func (f *AndroidFactory) CreateLabel() components.Label {
	return &components.AndroidLabel{}
}

func (f *AndroidFactory) CreateWindow() components.Window {
	return &components.AndroidWindow{}
}

type IOSFactory struct {
}

func (f *IOSFactory) CreateButton() components.Button {
	return &components.IOSButton{}
}

func (f *IOSFactory) CreateLabel() components.Label {
	return &components.IOSLabel{}
}

func (f *IOSFactory) CreateWindow() components.Window {
	return &components.IOSWindow{}
}

func ShowUI(factory Factory) {
	button := factory.CreateButton()
	label := factory.CreateLabel()
	window := factory.CreateWindow()
	button.Render()
	label.Render()
	window.Render()
}
