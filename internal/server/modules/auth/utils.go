package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SafeUser struct {
	ID             uuid.UUID
	Username       string
	Email          string
	Bio            *string
	Image          string
	FollowersCount int64
	FollowingCount int64
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

func (am *AuthModule) generateUserToken(user SafeUser) (string, error) {
	var hmacSampleSecret = []byte(am.cfg.Auth().Secret())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.ID, "email": user.Email})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
