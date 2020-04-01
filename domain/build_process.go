package domain

//  Builder Process for the Electronic Product
// the first three methods return the BuildProcess interface itself.
// it will help us to create the step by step of the creation object.
type BuildProcess interface {
	SetStructure() BuildProcess
	SetCamera() BuildProcess
	SetMonitor() BuildProcess
	GetGadget() ElectronicProduct
}
