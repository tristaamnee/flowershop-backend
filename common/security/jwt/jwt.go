package jwt

import (
	"github.com/go-pg/pg/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tristaamnee/flowershop-be/common/utils/configuration"
	tokenModel "github.com/tristaamnee/flowershop-be/tokens/model"
	tokenRepo "github.com/tristaamnee/flowershop-be/tokens/repository"
	"github.com/tristaamnee/flowershop-be/users/model"
	"time"
)

func GenerateToken(db *pg.DB, user *model.User) (string, error) {
	tokenID := uuid.New()

	//Define the JWT claims
	claims := jwt.MapClaims{
		"sub":    user.Email,
		"access": user.Role,
		"id":     tokenID,
		"exp":    time.Now().Add(time.Hour).Unix(), //Token expires in 1 hour
	}

	//CreateToken the JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtConfig, _ := configuration.Get("jwt")

	tokenString, err := token.SignedString([]byte(jwtConfig.(configuration.Jwt).EncryptionKey))
	if err != nil {
		return "", err
	}

	err = persistToken(db, user, tokenID, err)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func persistToken(db *pg.DB, user *model.User, tokenID uuid.UUID, err error) error {

	tokenStruct := tokenModel.Token{
		ID:           tokenID,
		CreationDate: time.Now(),
		UserID:       user.ID,
	}

	err = tokenRepo.CreateToken(db, &tokenStruct)
	return err
}
