package gaspool

type Gas uint64

func InitGas(gas uint64) Gas {
	return Gas(gas)
}

func (g *Gas) AddGas(gas uint64) {
	*g += Gas(gas)
}

func (g *Gas) SubGas(gas uint64) {
	*g -= Gas(gas)
}

func (g *Gas) SetGas(gas uint64) {
	*g = Gas(gas)
}
