package interfaces

type Shipper interface {
	ShipperName() string
	ShipperAddress() string
	ShipperSubDistrice() string
	ShipperDistrice() string
	ShipperProvince() string
	ShipperZipCode() string
	ShipperPhone() string
	ShipperLat() string
	ShipperLng() string
}