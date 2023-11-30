package dto

import "aseprayana-skeleton-go/dto"

type CreateRequest struct {
	Goods []dto.GoodsRequest `json:"goods"`
}

type GetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type DeleteRequest struct {
	ID string `param:"id" validate:"required"`
}

type UpdateRequest struct {
	ID            string `param:"id" validate:"required"`
	NomorKontrak  string `json:"nomor_kontrak" form:"nomor_kontrak" validate:"required,nomor_kontrak"`
	OTR           string `json:"otr" form:"otr" validate:"required,otr"`
	AdminFee      string `json:"admin_fee"`
	JumlahCicilan string `json:"jumlah_cicilan"`
	JumlahBunga   string `json:"jumlah_bunga"`
	NamaAsset     string `json:"nama_asset"`
}

type TransactionResponse struct {
	NomorKontrak  string `json:"nomor_kontrak" form:"nomor_kontrak" validate:"required,nomor_kontrak"`
	OTR           string `json:"otr" form:"otr" validate:"required,otr"`
	AdminFee      string `json:"admin_fee"`
	JumlahCicilan string `json:"jumlah_cicilan"`
	JumlahBunga   string `json:"jumlah_bunga"`
	NamaAsset     string `json:"nama_asset"`
}
