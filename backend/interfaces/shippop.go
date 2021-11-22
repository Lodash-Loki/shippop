package interfaces

type Shippop interface {
	Ondemand()
	Dropoff()
	Dropship()
	Frozen()
}