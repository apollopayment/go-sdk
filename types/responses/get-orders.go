package responses

type GetOrders struct {
	Orders []Order `json:"orders"`
}
