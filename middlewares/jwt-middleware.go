package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(id uint, name string) (string, error)

	ValidateToken(token string) (*jwt.Token, string, error)
}

type jwtCustomClaim struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "admin",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "system"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(id uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"name": name,
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
		"iss":  j.issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}
func (j *jwtService) ValidateToken(token string) (*jwt.Token, string, error) {
	// Trim leading and trailing whitespaces
	token = strings.TrimSpace(token)
	fmt.Println("Debug: Token to be validated:", token)

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return nil, "", errors.New("Invalid token format")
	}

	parsedToken, err := jwt.Parse(splitToken[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		fmt.Println("Debug: Error parsing token:", err)
		return nil, "", err
	}

	// Extract the user ID from the claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Debug: Failed to extract claims from token")
		return nil, "", errors.New("Failed to extract claims from token")
	}

	userID, ok := claims["name"].(string)
	if !ok {
		fmt.Println("Debug: 'name' not found in JWT claims")
		return nil, "", errors.New("'name' not found in JWT claims")
	}

	return parsedToken, userID, nil
}

// func (j *jwtService) ValidateToken(token string) (*jwt.Token, string, error) {

// 	splitToken := strings.Split(token, "Bearer ")
// 	if len(splitToken) != 1 {
// 		return nil, "", errors.New("Invalid token format")
// 	}

// 	t, err := jwt.Parse(splitToken[1], func(t_ *jwt.Token) (interface{}, error) {
// 		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
// 		}
// 		return []byte(j.secretKey), nil
// 	})

// 	if err != nil {
// 		return nil, "", err
// 	}

// 	// Extract the user ID from the claims
// 	claims, ok := t.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, "", errors.New("Failed to extract claims from token")
// 	}

// 	userID, ok := claims["name"].(string)
// 	if !ok {
// 		return nil, "", errors.New("name not found in JWT claims")
// 	}

// 	return t, userID, nil
// }
