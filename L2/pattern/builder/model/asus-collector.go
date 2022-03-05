package model

type AsusCollector struct {
	Core  int
	Brand string
	RAM   int
	GPU   int
}

func (c *AsusCollector) SetCore() {
	c.Core = 8
}

func (c *AsusCollector) SetBrand() {
	c.Brand = "ASUS"
}

func (c *AsusCollector) SetRAM() {
	c.RAM = 8
}

func (c *AsusCollector) SetGPU() {
	c.GPU = 1
}

func (c *AsusCollector) GetPC() PC {
	return PC{
		Core:  c.Core,
		Brand: c.Brand,
		RAM:   c.RAM,
		GPU:   c.GPU,
	}
}
