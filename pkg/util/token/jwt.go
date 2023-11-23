package token

import (
	"time"

	"github.com/arvinpaundra/ngekost-api/pkg/util/config"
	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type JWTCustomClaim struct {
	UserId    string
	Role      string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

type JSONWebToken interface {
	Encode(claims *JWTCustomClaim) (string, error)
	Decode(token string) (*JWTCustomClaim, error)
}

type jsonWebToken struct {
	secret string
}

func NewJWT() JSONWebToken {
	return &jsonWebToken{
		secret: config.GetString("JWT_SECRET"),
	}
}

func (j *jsonWebToken) Encode(claims *JWTCustomClaim) (string, error) {
	c := jwtClaim{
		UserId: claims.UserId,
		Role:   claims.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  claims.IssuedAt.Unix(),
			ExpiresAt: claims.ExpiredAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(j.secret))
}

func (j *jsonWebToken) Decode(token string) (*JWTCustomClaim, error) {
	panic("not implemented")
}
