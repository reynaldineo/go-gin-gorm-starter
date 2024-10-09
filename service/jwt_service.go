package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/reynaldineo/go-gin-gorm-starter/dto"
)

type (
	JWTService interface {
		GenerateToken(userId string, role string) string
		ValidateToken(token string) (*jwt.Token, error)
		GetPayloadInsideToken(token string) (string, string, error)
	}

	jwtService struct {
		secretKey string
		issuer    string
	}
)

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "issuer_reynaldineo",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "secretKey_reynaldineo"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userId string, role string) string {
	claims := jwtCustomClaim{
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60 * 2)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}

	return tokenString
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}

func (j *jwtService) GetPayloadInsideToken(token string) (string, string, error) {
	t_Token, err := j.ValidateToken(token)
	if err != nil {
		return "", "", err
	}

	if !t_Token.Valid {
		return "", "", dto.ErrTokenInvalid
	}

	claims := t_Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	role := fmt.Sprintf("%v", claims["role"])

	return id, role, nil
}
