package interfaces

type Parcel interface {
	Weight() float64
	Size() string
	Detail() string
	Price() float64
	Category() string
}