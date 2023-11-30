package middlewares

import (
	res "aseprayana-skeleton-go/util/response"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthorizeJWT(jwtService JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return res.ErrorBuilder(&res.ErrorConstant.BadRequest, &res.Error{}).Send(c)
			}

			// ValidateToken returns the token, userID, and an error
			token, _, err := jwtService.ValidateToken(authHeader)
			if err != nil {
				// Handle the error, e.g., log it or return an internal server error
				return err
			}

			// Now you have the token and userID for further processing
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println("Claim[name]: ", claims["name"])
				return next(c)
			} else {
				return res.ErrorBuilder(&res.ErrorConstant.BadRequest, &res.Error{}).Send(c)
			}
		}
	}
}
