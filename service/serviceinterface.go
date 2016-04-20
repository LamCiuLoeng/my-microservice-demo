package main

type ServiceInterface interface {
	Register(name, version, node string) error
	Destroy(name, version, node string) error
	Send() error
	Receive() error
}

type RegistryInterface interface {
	RegisterService(name, version, node string) error
	DestroyService(name, version, node string) error
	ListServices() error
	GetService() (ServiceInterface, error)
}
