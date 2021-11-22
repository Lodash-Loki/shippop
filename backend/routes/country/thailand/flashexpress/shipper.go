package flashexpress

import (
	"github.com/Lodash-Loki/shippop/domain"
)

var Shipper shipper

type shipper struct{}

var shipperModel *domain.Shipper

func (s *shipper) ShipperName() string {
	return "Name" 
}

func (s *shipper) ShipperSubDistrice() string {
	return "SubDistrice"
}

func (s *shipper) ShipperDistrice() string {
	return "Districe"
}

func (s *shipper) ShipperProvince() string {
	return "Province"
}

func (s *shipper) ShipperZipCode() string {
	return "ZipCode"
}

func (s *shipper) ShipperPhone() string {
	return "Phone"
}

func (s *shipper) ShipperLat() string {
	return "Lat"
}

func (s *shipper) ShipperLng() string {
	return "Lng"
}
