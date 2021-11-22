package interfaces

type Customer interface {
	Name() string
	Address() string
	SubDistrice() string
	Districe() string
	Province() string
	ZipCode() string
	Phone() string
	Lat() string
	Lng() string
}