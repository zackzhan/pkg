package text

//go:generate stringer -type=OrderStatus
type OrderStatus int

const (
	CREATE OrderStatus = iota + 1
	PAID
	DELIVERING
	COMPLETED
	CANCELLED
)
