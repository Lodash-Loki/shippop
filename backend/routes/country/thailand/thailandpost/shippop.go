package flashexpress

func (s *shipper) Ondemand() string {
	return "Ondemand"
}

func (s *shipper) Dropoff() string {
	return "Dropoff"
}

func (s *shipper) Dropship() string {
	return "Dropship"
}

func (s *shipper) Frozen() string {
	return "Frozen"
}
