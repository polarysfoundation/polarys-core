package gaspool

type GasPool struct {
	lowBase  Gas
	midBase  Gas
	highBase Gas
}

// InitGasPool initializes a GasPool with calculated gas values based on the baseGas input.
func InitGasPool(baseGas uint64) *GasPool {
	return &GasPool{
		lowBase:  calculateLowGas(baseGas),
		midBase:  calculateMidGas(baseGas),
		highBase: calculateHighGas(baseGas),
	}
}

func (gp *GasPool) Low() Gas  { return gp.lowBase }
func (gp *GasPool) Base() Gas { return gp.midBase }
func (gp *GasPool) High() Gas { return gp.highBase }

func (gp *GasPool) Update(base uint64) {
	gp.lowBase = calculateLowGas(base)
	gp.midBase = calculateMidGas(base)
	gp.highBase = calculateHighGas(base)

}

// calculateLowGas calculates the low base gas value.
func calculateLowGas(baseGas uint64) Gas {
	return Gas((baseGas * 2) / 3)
}

// calculateMidGas calculates the mid base gas value.
func calculateMidGas(baseGas uint64) Gas {
	return Gas((baseGas * 4) / 5)
}

// calculateHighGas calculates the high base gas value.
func calculateHighGas(baseGas uint64) Gas {
	return Gas((baseGas * 6) / 5)
}
