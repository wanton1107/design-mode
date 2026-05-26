package adapter

import "fmt"

type Target interface {
	Request() string
}

type XMLSystem struct {
}

func (x *XMLSystem) Request() string {
	return "这里是XML格式的数据"
}

type JSONSystem struct {
	*XMLSystem
}

func (j *JSONSystem) Request() string {
	xmlData := j.XMLSystem.Request()
	fmt.Println("XML数据转换为JSON数据", xmlData)
	return "这里是JSON格式的数据"
}

