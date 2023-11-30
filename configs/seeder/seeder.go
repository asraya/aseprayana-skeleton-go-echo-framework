package seeder

import (
	"aseprayana-skeleton-go/configs"
	"aseprayana-skeleton-go/entity"
	"fmt"
)

func Users() {
	db := configs.InitDB()
	// now := time.Now()

	var users []entity.User
	var budi = entity.User{
		ID:           1,
		FullName:     "Budi",
		LegalName:    "Budi",
		Email:        "budi@gmail.com",
		TempatLahir:  "jakarta",
		TanggalLahir: "1/2/1995",
		Gaji:         "5.000.000",
		FotoKtp:      "jpg",
		FotoSelfie:   "jpg",
		Password:     "$2a$10$k7oCB0eh840JtXnQ74OnUezmBuYLQnmdbXOLc3ztN9F7y/C8jHFE6", //12345678

	}
	users = append(users, budi)

	var annisa = entity.User{
		ID:           2,
		FullName:     "Annisa",
		LegalName:    "Annisa",
		Email:        "annisa@gmail.com",
		TempatLahir:  "jakarta",
		TanggalLahir: "3/4/1995",
		Gaji:         "10.000.000",
		FotoKtp:      "jpg",
		FotoSelfie:   "jpg",
		Password:     "$2a$10$k7oCB0eh840JtXnQ74OnUezmBuYLQnmdbXOLc3ztN9F7y/C8jHFE6", //12345678

	}
	users = append(users, annisa)

	for _, Users := range users {
		if err := db.Create(&Users).Error; err != nil {
			db.Create(&Users)
		}
		fmt.Printf("package type %s has been created\n", Users.Email)

	}
}

func Limit() {
	db := configs.InitDB()

	var limits []entity.Limit

	var limit = entity.Limit{
		ID:     1,
		UserID: 1,
		Limit:  100000,
		Month:  1,
	}
	limits = append(limits, limit)

	var limit2 = entity.Limit{
		ID:     2,
		UserID: 1,
		Limit:  200000,
		Month:  2,
	}
	limits = append(limits, limit2)

	var limit3 = entity.Limit{
		ID:     3,
		UserID: 1,
		Limit:  500000,
		Month:  3,
	}
	limits = append(limits, limit3)

	var limit4 = entity.Limit{
		ID:     4,
		UserID: 1,
		Limit:  700000,
		Month:  4,
	}
	limits = append(limits, limit4)

	var limit5 = entity.Limit{
		ID:     5,
		UserID: 2,
		Limit:  1000000,
		Month:  1,
	}
	limits = append(limits, limit5)

	var limit6 = entity.Limit{
		ID:     6,
		UserID: 2,
		Limit:  1200000,
		Month:  2,
	}
	limits = append(limits, limit6)

	var limit7 = entity.Limit{
		ID:     7,
		UserID: 2,
		Limit:  1500000,
		Month:  3,
	}
	limits = append(limits, limit7)

	var limit8 = entity.Limit{
		ID:     8,
		UserID: 2,
		Limit:  2000000,
		Month:  4,
	}
	limits = append(limits, limit8)
	for _, Limits := range limits {
		if err := db.Create(&Limits).Error; err != nil {
			db.Create(&Limits)
		}

	}
}

func Goods() {
	db := configs.InitDB()

	var gds []entity.Goods

	var gods1 = entity.Goods{
		ID:    1,
		Price: 300000,
		Goods: "mouse",
	}
	gds = append(gds, gods1)

	var gods2 = entity.Goods{
		ID:    2,
		Price: 100000,
		Goods: "tas",
	}
	gds = append(gds, gods2)

	for _, Goods := range gds {
		if err := db.Create(&Goods).Error; err != nil {
			db.Create(&Goods)
		}
		fmt.Printf("package type %s has been created\n", Goods.Goods)

	}
}

func RunSeeder() {
	Users()
	Limit()
	Goods()
}
