package dto

import (
	"aseprayana-skeleton-go/dto"
)

type CreateRequest struct {
	CreatedBy string             `json:"created_by"`
	Goods     []dto.GoodsRequest `json:"goods"`
}

type GetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type DeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type UpdateRequest struct {
	ID            string `param:"id" validate:"required"`
	NomorKontrak  string `json:"nomor_kontrak" form:"nomor_kontrak" validate:"required,nomor_kontrak"`
	OTR           string `json:"otr" form:"otr" validate:"required,otr"`
	AdminFee      string `json:"admin_fee"`
	JumlahCicilan string `json:"jumlah_cicilan"`
	JumlahBunga   string `json:"jumlah_bunga"`
	NamaAsset     string `json:"nama_asset"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	DeletedBy     string `json:"deleted_by"`
}

type TransactionResponse struct {
	ID            string `json:"id"`
	NomorKontrak  string `json:"nomor_kontrak" form:"nomor_kontrak" validate:"required,nomor_kontrak"`
	OTR           string `json:"otr" form:"otr" validate:"required,otr"`
	AdminFee      string `json:"admin_fee"`
	JumlahCicilan string `json:"jumlah_cicilan"`
	JumlahBunga   string `json:"jumlah_bunga"`
	NamaAsset     string `json:"nama_asset"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	DeletedBy     string `json:"deleted_by"`
}
