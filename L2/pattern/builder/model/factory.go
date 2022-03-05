package model

type Factory struct {
	collector Collector
}

func NewFactory(c Collector) *Factory {
	return &Factory{collector: c}
}

func (f *Factory) ChangeCollector(c Collector) {
	f.collector = c
}

func (f Factory) CreateComputer() PC {
	f.collector.SetCore()
	f.collector.SetBrand()
	f.collector.SetGPU()
	f.collector.SetRAM()

	return f.collector.GetPC()
}
