package domain

type Inventory struct {
	ProductID         string
	Quantity          int
	LowStockThreshold int
}
