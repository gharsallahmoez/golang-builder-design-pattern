package domain


type Laptop struct {
	electronicProduct ElectronicProduct
}

func (l *Laptop) SetStructure() BuildProcess {
	l.electronicProduct.Structure = "Laptop"
	return l
}

func (l *Laptop) SetMonitor() BuildProcess {
	l.electronicProduct.Monitor = 1
	return l
}

func (l *Laptop) SetCamera() BuildProcess {
	l.electronicProduct.Camera = 1
	return l
}

func (l *Laptop) GetGadget() ElectronicProduct {
	return l.electronicProduct
}