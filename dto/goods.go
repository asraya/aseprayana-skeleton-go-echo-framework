package dto

type GoodsRequest struct {
	Price float64 `json:"price"`
	Tenor int     `json:"tenor"`
	Goods string  `json:"goods"`
}
