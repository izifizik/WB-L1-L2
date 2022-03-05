package model

type Collector interface {
	SetCore()
	SetBrand()
	SetRAM()
	SetGPU()
	GetPC() PC
}

func GetCollector(c string) Collector {
	switch c {
	case "ASUS":
		return &AsusCollector{}
	case "HP":
		return &HPCollector{}
	default:
		return nil
	}
}
