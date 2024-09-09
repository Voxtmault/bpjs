package services

import "github.com/voxtmault/bpjs-rs-module/pkg/interfaces"

type DependencyInjectionService struct {
	// You can modify this struct as needed
}

var _ interfaces.DependencyInjection = &DependencyInjectionService{}

func (di *DependencyInjectionService) InjectionSample() error {
	// Do Something
	return nil
}
