package main

import (
	"design-mode/adapter"
	"fmt"
)

func main() {
	xmlSystem := &adapter.XMLSystem{}
	jsonSystem := &adapter.JSONSystem{XMLSystem: xmlSystem}
	fmt.Println(jsonSystem.Request())
}
