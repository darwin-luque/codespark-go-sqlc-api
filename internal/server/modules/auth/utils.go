package auth

import (
	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/golang-jwt/jwt"
)

func (am *AuthModule) generateUserToken(user *domain.CreateUserRow) (string, error) {
	var hmacSampleSecret = []byte(am.cfg.Auth().Secret())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.ID, "email": user.Email})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
