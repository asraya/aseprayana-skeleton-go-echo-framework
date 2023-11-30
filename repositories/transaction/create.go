package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"aseprayana-skeleton-go/entity"
	"aseprayana-skeleton-go/util"

	"strconv"
	"strings"
)

func (b *transactionRepository) Create(req dto.CreateRequest) (*dto.TransactionResponse, error) {
	var gds []entity.Goods
	var transaction []entity.Transaction

	for _, goodsReq := range req.Goods {
		gds = append(gds, entity.Goods{
			Goods: goodsReq.Goods,
			Price: goodsReq.Price,
			Tenor: goodsReq.Tenor,
		})
	}

	for _, goods := range gds {
		transaction = append(transaction, entity.Transaction{
			NamaAsset: goods.Goods,
		})
	}

	if err := b.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	var namaAssets []string
	for _, goods := range gds {
		namaAssets = append(namaAssets, goods.Goods)
	}

	joinedNames := strings.Join(namaAssets, ", ")
	adminFee := 3000.0
	jumlahCicilanStr := util.CalculateJumlahCicilan(gds, adminFee, 2)

	// Convert JumlahCicilan to float64
	jumlahCicilan, err := strconv.ParseFloat(jumlahCicilanStr, 64)
	if err != nil {
		return nil, err
	}
	// Assuming you have a Tenor field in your entity.Transaction struct
	tenor := gds[0].Tenor
	harga := gds[0].Price

	// Assuming you have a Rate field in your entity.Transaction struct
	rate := 0.05 // replace with your actual Rate logic if needed

	// Calculate JumlahBunga using the calculateJumlahBunga function
	jumlahBunga := util.CalculateJumlahBunga(jumlahCicilan, rate, tenor)
	jumlahCicilanTotal := jumlahBunga + harga + adminFee

	response := &dto.TransactionResponse{

		NomorKontrak:  util.GenerateRandomString(),
		OTR:           util.GenerateOTR(8),
		AdminFee:      util.FormatIDR(adminFee),
		JumlahCicilan: util.FormatIDR(jumlahCicilanTotal),
		JumlahBunga:   util.FormatIDR(jumlahBunga),
		NamaAsset:     joinedNames,
	}

	if err := b.DB.Create(&entity.Transaction{
		NamaAsset:     response.NamaAsset,
		NomorKontrak:  response.NomorKontrak,
		OTR:           response.OTR,
		AdminFee:      response.AdminFee,
		JumlahCicilan: response.JumlahCicilan,
		JumlahBunga:   response.JumlahBunga,
	}).Error; err != nil {
		return nil, err
	}

	return response, nil
}
