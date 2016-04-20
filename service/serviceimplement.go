package main

type HTTPService struct {
	Name    string
	Version string
	Host    string
	Port    string
}

func (s HTTPService) Register(name, version, node) error {
	return nil
}
