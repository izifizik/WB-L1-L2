package data

import "fmt"

type JsonDoc struct {
	data string
}

type JsonDocAdapter struct {
	jDoc JsonDoc
}

func NewJsonDoc(data string) JsonDoc {
	return JsonDoc{data: data}
}

func NewJsonDocAdapter(doc JsonDoc) JsonDocAdapter {
	return JsonDocAdapter{jDoc: doc}
}

func (d JsonDoc) ConvertToXML() string {
	return "<xml " + d.data + " />"
}

func (d JsonDocAdapter) PrintDoc() {
	fmt.Println("The doc: (convert)" + d.jDoc.ConvertToXML())
}
