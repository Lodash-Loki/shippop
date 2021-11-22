package flashexpress

var Merchant merchant

type merchant struct{}

func (m *merchant) MechantName() string {
	return "Name"
}

func (m *merchant) MechantSubDistrice() string {
	return "SubDistrice"
}

func (m *merchant) MechantDistrice() string {
	return "Districe"
}

func (m *merchant) MechantProvince() string {
	return "Province"
}

func (m *merchant) MechantZipCode() string {
	return "ZipCode"
}

func (m *merchant) MechantPhone() string {
	return "Phone"
}

func (m *merchant) MechantLat() string {
	return "Lat"
}

func (m *merchant) MechantLng() string {
	return "Lng"
}
