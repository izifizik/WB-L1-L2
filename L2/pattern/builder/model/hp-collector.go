package model

type HPCollector struct {
	Core  int
	Brand string
	RAM   int
	GPU   int
}

func (c *HPCollector) SetCore() {
	c.Core = 16
}

func (c *HPCollector) SetBrand() {
	c.Brand = "HP"
}

func (c *HPCollector) SetRAM() {
	c.RAM = 32
}

func (c *HPCollector) SetGPU() {
	c.GPU = 2
}

func (c *HPCollector) GetPC() PC {
	return PC{
		Core:  c.Core,
		Brand: c.Brand,
		RAM:   c.RAM,
		GPU:   c.GPU,
	}
}
