package util

import (
	"aseprayana-skeleton-go/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Negotiate(c echo.Context, code int, i interface{}) error {
	mediaType := c.QueryParam("mediaType")

	switch mediaType {
	case "xml":
		return c.XML(code, i)
	case "json":
		return c.JSON(code, i)
	default:
		return c.JSON(code, i)
	}
}

func IsDuplicateEntryError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return mysqlErr.Number == 1062
	}
	return false
}

func GenerateRandomString() string {
	randomUUID := uuid.New()
	return randomUUID.String()
}

func GenerateOTR(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func CalculateJumlahCicilan(goods []entity.Goods, adminFee float64, precision int) string {
	var total float64

	for _, g := range goods {
		total += g.Price
	}

	total += adminFee

	return fmt.Sprintf("%.*f", precision, total)
}

func FormatIDR(value float64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("%s", currency.IDR.Amount(value))
}

func CalculateJumlahBunga(principal float64, rate float64, time int) float64 {
	// Formula: Bunga = Principal * Rate * Time
	bunga := principal * rate * float64(time)
	return bunga
}
