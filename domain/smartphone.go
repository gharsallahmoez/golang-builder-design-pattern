package domain

type Smartphone struct {
	electronicProduct ElectronicProduct
}

func (s *Smartphone) SetStructure() BuildProcess {
	s.electronicProduct.Structure = "Smartphone"
	return s
}

func (s *Smartphone) SetMonitor() BuildProcess {
	s.electronicProduct.Monitor = 1
	return s
}

func (s *Smartphone) SetCamera() BuildProcess {
	s.electronicProduct.Camera = 2
	return s
}

func (s *Smartphone) GetGadget() ElectronicProduct {
	return s.electronicProduct
}