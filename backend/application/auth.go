package application

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func (app *Application) CreateToken(wallet string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": wallet,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
			"iss": "localhost",
		})

	apikey, err := app.DB.AdminGetWalletAccount(wallet)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(apikey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (app *Application) verifyToken(tokenString, wallet string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		apikey, err := app.DB.AdminGetWalletAccount(wallet)
		if err != nil {
			return "", err
		}
		return apikey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
