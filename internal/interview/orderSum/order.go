package orderSum

type order struct {
	timestamp    string
	postalCode   string
	customerName string
	unitVolume   int
}

type orderList []order
