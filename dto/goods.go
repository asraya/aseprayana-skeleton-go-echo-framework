package dto

type GoodsRequest struct {
	Price     float64 `json:"price"`
	Tenor     int     `json:"tenor"`
	Goods     string  `json:"goods"`
	CreatedBy string  `json:"created_by"`
	UpdatedBy string  `json:"updated_by"`
	DeletedBy string  `json:"deleted_by"`
}
