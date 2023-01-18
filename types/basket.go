package types

type AddtoBasketRequest struct {
	Sku   string  `json:"sku"`
	Price float32 `json:"price"`
}
