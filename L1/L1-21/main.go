package main

import (
	"Work/L1/L1-21/data"
	"Work/L1/L1-21/service"
)

func main() {
	d := "XML data"
	xml := service.NewXMLDoc(d)
	xml.PrintDoc()

	d = "json data"
	json := data.NewJsonDoc(d)
	jAdapter := data.NewJsonDocAdapter(json)
	jAdapter.PrintDoc()
}
