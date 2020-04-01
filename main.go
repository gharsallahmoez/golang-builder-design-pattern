package main

import "github.com/gharsallahmoez/golang-builder-design-pattern/domain"

func main() {
	manufacturingDirector := domain.ManufacturingDirector{}

	laptop := &domain.Laptop{}
	manufacturingDirector.SetBuilder(laptop)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()

	smartphone := &domain.Smartphone{}
	manufacturingDirector.SetBuilder(smartphone)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
}