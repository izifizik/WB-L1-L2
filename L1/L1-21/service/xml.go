package service

import "fmt"

type XMLService interface {
	PrintDoc()
}

type XMLDoc struct {
	data string
}

func NewXMLDoc(data string) XMLService {
	return &XMLDoc{data: data}
}

func (d XMLDoc) PrintDoc() {
	fmt.Println("The doc: <xml " + d.data + " />")
}
