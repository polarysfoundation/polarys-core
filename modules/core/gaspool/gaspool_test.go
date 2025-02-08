package gaspool

import (
	"log"
	"testing"
)

func TestInitGasPool(t *testing.T) {
	baseGas := uint64(100)
	log.Printf("Initializing GasPool with baseGas: %v", baseGas)
	gp := InitGasPool(baseGas)

	if gp.lowBase != calculateLowGas(baseGas) {
		t.Errorf("Expected lowBase %v, got %v", calculateLowGas(baseGas), gp.lowBase)
	}
	if gp.midBase != calculateMidGas(baseGas) {
		t.Errorf("Expected midBase %v, got %v", calculateMidGas(baseGas), gp.midBase)
	}
	if gp.highBase != calculateHighGas(baseGas) {
		t.Errorf("Expected highBase %v, got %v", calculateHighGas(baseGas), gp.highBase)
	}
}

func TestGasPool_Update(t *testing.T) {
	baseGas := uint64(100)
	log.Printf("Initializing GasPool with baseGas: %v", baseGas)
	gp := InitGasPool(baseGas)

	prevLowBase := gp.lowBase
	prevBase := gp.midBase
	prevHighBase := gp.highBase

	log.Printf("LowGas: %d", prevLowBase)
	log.Printf("BaseGas: %d", prevBase)
	log.Printf("HighGas: %d", prevHighBase)

	newBase := uint64(50)
	log.Printf("Updating GasPool with newBase: %v", newBase)
	gp.Update(newBase)

	if gp.lowBase != calculateLowGas(uint64(newBase)) {
		t.Errorf("Expected lowBase %v, got %v", calculateLowGas(uint64(newBase)), gp.lowBase)
	}
	if gp.midBase != calculateMidGas(newBase) {
		t.Errorf("Expected midBase %v, got %v", calculateMidGas(newBase), gp.midBase)
	}
	if gp.highBase != calculateHighGas(uint64(newBase)) {
		t.Errorf("Expected highBase %v, got %v", calculateHighGas(uint64(newBase)), gp.highBase)
	}

	prevLowBase = gp.lowBase
	prevBase = gp.midBase
	prevHighBase = gp.highBase

	log.Printf("LowGas: %d", prevLowBase)
	log.Printf("BaseGas: %d", prevBase)
	log.Printf("HighGas: %d", prevHighBase)

	newBase = uint64(200)
	log.Printf("Updating GasPool with newBase: %v", newBase)
	gp.Update(newBase)

	if gp.lowBase != calculateLowGas(uint64(newBase)) {
		t.Errorf("Expected lowBase %v, got %v", calculateLowGas(uint64(newBase)), gp.lowBase)
	}
	if gp.midBase != calculateMidGas(newBase) {
		t.Errorf("Expected midBase %v, got %v", calculateMidGas(newBase), gp.midBase)
	}
	if gp.highBase != calculateHighGas(uint64(newBase)) {
		t.Errorf("Expected highBase %v, got %v", calculateHighGas(uint64(newBase)), gp.highBase)
	}

	prevLowBase = gp.lowBase
	prevBase = gp.midBase
	prevHighBase = gp.highBase

	log.Printf("LowGas: %d", prevLowBase)
	log.Printf("BaseGas: %d", prevBase)
	log.Printf("HighGas: %d", prevHighBase)

	newBase = uint64(1000)
	log.Printf("Updating GasPool with newBase: %v", newBase)
	gp.Update(newBase)

	if gp.lowBase != calculateLowGas(uint64(newBase)) {
		t.Errorf("Expected lowBase %v, got %v", calculateLowGas(uint64(newBase)), gp.lowBase)
	}
	if gp.midBase != calculateMidGas(newBase) {
		t.Errorf("Expected midBase %v, got %v", calculateMidGas(newBase), gp.midBase)
	}
	if gp.highBase != calculateHighGas(uint64(newBase)) {
		t.Errorf("Expected highBase %v, got %v", calculateHighGas(uint64(newBase)), gp.highBase)
	}

	prevLowBase = gp.lowBase
	prevBase = gp.midBase
	prevHighBase = gp.highBase

	log.Printf("LowGas: %d", prevLowBase)
	log.Printf("BaseGas: %d", prevBase)
	log.Printf("HighGas: %d", prevHighBase)

	newBase = uint64(10)
	log.Printf("Updating GasPool with newBase: %v", newBase)
	gp.Update(newBase)

	if gp.lowBase != calculateLowGas(uint64(newBase)) {
		t.Errorf("Expected lowBase %v, got %v", calculateLowGas(uint64(newBase)), gp.lowBase)
	}
	if gp.midBase != calculateMidGas(newBase) {
		t.Errorf("Expected midBase %v, got %v", calculateMidGas(newBase), gp.midBase)
	}
	if gp.highBase != calculateHighGas(uint64(newBase)) {
		t.Errorf("Expected highBase %v, got %v", calculateHighGas(uint64(newBase)), gp.highBase)
	}

	prevLowBase = gp.lowBase
	prevBase = gp.midBase
	prevHighBase = gp.highBase

	log.Printf("LowGas: %d", prevLowBase)
	log.Printf("BaseGas: %d", prevBase)
	log.Printf("HighGas: %d", prevHighBase)

	newBase = uint64(13)
	log.Printf("Updating GasPool with newBase: %v", newBase)
	gp.Update(newBase)

	if gp.lowBase != calculateLowGas(uint64(newBase)) {
		t.Errorf("Expected lowBase %v, got %v", calculateLowGas(uint64(newBase)), gp.lowBase)
	}
	if gp.midBase != calculateMidGas(newBase) {
		t.Errorf("Expected midBase %v, got %v", calculateMidGas(newBase), gp.midBase)
	}
	if gp.highBase != calculateHighGas(uint64(newBase)) {
		t.Errorf("Expected highBase %v, got %v", calculateHighGas(uint64(newBase)), gp.highBase)
	}

	prevLowBase = gp.lowBase
	prevBase = gp.midBase
	prevHighBase = gp.highBase

	log.Printf("LowGas: %d", prevLowBase)
	log.Printf("BaseGas: %d", prevBase)
	log.Printf("HighGas: %d", prevHighBase)
}

func TestGasPool_Low(t *testing.T) {
	baseGas := uint64(100)
	log.Printf("Initializing GasPool with baseGas: %v", baseGas)
	gp := InitGasPool(baseGas)

	if gp.Low() != calculateLowGas(baseGas) {
		t.Errorf("Expected Low %v, got %v", calculateLowGas(baseGas), gp.Low())
	}
}

func TestGasPool_Base(t *testing.T) {
	baseGas := uint64(100)
	log.Printf("Initializing GasPool with baseGas: %v", baseGas)
	gp := InitGasPool(baseGas)

	if gp.Base() != calculateMidGas(baseGas) {
		t.Errorf("Expected Base %v, got %v", calculateMidGas(baseGas), gp.Base())
	}
}

func TestGasPool_High(t *testing.T) {
	baseGas := uint64(100)
	log.Printf("Initializing GasPool with baseGas: %v", baseGas)
	gp := InitGasPool(baseGas)

	if gp.High() != calculateHighGas(baseGas) {
		t.Errorf("Expected High %v, got %v", calculateHighGas(baseGas), gp.High())
	}
}
