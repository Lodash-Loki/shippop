package interfaces

type Merchant interface {
	MerchantName() string
	MerchantAddress() string
	MerchantSubDistrice() string
	MerchantDistrice() string
	MerchantProvince() string
	MerchantZipCode() string
	MerchantPhone() string
	MerchantLat() string
	MerchantLng() string
	MerchantZipCodeDrop() string
}