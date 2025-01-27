package auth

import (
	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (am *AuthModule) generateUserToken(user *domain.User) (string, error) {
	var hmacSampleSecret = []byte(am.cfg.Auth().Secret())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.ID, "email": user.Email})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type SafeUser struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Bio       string
	Image     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
