package composite

import "fmt"
type Component interface {
	Operation()
}

type File struct {
	Name string
}

func (f *File) Operation() {
	fmt.Println("File:", f.Name)
}

type Folder struct {
	Name string
	Children []Component
}

func (f *Folder) Operation() {
	fmt.Println("Folder:", f.Name)
	for _, child := range f.Children {
		child.Operation()
	}
}

func (f *Folder) Add(component Component) {
	f.Children = append(f.Children, component)
}

func (f *Folder) Remove(component Component) {
	for i, child := range f.Children {
		if child == component {
			f.Children = append(f.Children[:i], f.Children[i+1:]...)
		}
	}
}
