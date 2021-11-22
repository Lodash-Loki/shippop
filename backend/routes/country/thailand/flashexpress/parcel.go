package flashexpress

import "github.com/Lodash-Loki/shippop/domain"

var Parcel parcel

type parcel struct{}

var parcelModel *domain.Parcel

func (p *parcel) Weight() string {
	return "Name"
}

func (p *parcel) Size() string {
	return "SubDistrice"
}

func (p *parcel) Detail() string {
	return "Districe"
}

func (p *parcel) Price() string {
	return "Province"
}

func (p *parcel) Category() string {
	return "ZipCode"
}
