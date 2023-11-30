package response

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

// func IsNumber(data string) error {
// 	_, err := strconv.Atoi(data)
// 	if err != nil {
// 		return fmt.Errorf("path parameter hanya menerima angka")
// 	}

// 	return nil
// }

func IsNumber(c echo.Context, paramName string) (uint, error) {
	value := c.Param(paramName)

	if len(value) == 0 {
		return 0, fmt.Errorf("%s cannot be empty", paramName)
	}

	uintValue, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("invalid format for %s: %v", paramName, err)
	}

	return uint(uintValue), nil
}
