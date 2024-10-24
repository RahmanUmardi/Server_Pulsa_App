package service

import (
	"fmt"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/shared/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	CreateToken(user entity.User) (dto.AuthResponseDto, error)
	ValidateToken(tokenString string) (*model.Claim, error)
}
type jwtService struct {
	cfgToken config.TokenConfig
}

func (j *jwtService) CreateToken(user entity.User) (dto.AuthResponseDto, error) {
	claims := model.Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfgToken.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfgToken.JwtExpiresTime)),
		},
		UserId: user.Id_user,
		Role:   user.Role,
	}

	token := jwt.NewWithClaims(j.cfgToken.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.cfgToken.JwtSignatureKy)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to create token: %v", err)
	}

	return dto.AuthResponseDto{Token: ss}, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*model.Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return j.cfgToken.JwtSignatureKy, nil
	})

	if err != nil {
		return nil, fmt.Errorf("unauthorized : %v", err)
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized : %v", err)
	}

	return claim, nil
}

func NewJwtService(cfgToken config.TokenConfig) JwtService {
	return &jwtService{cfgToken: cfgToken}
}
