package helpers

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iqquee/auth/database"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateTokens returns the access and refresh tokens
func GenerateTokens(uuid string) (string, string) {
	claim, accessToken := PostgresQlGenerateAccessClaims(uuid)
	refreshToken := PostgresQlGenerateRefreshClaims(claim)

	return accessToken, refreshToken
}

// GenerateAccessClaims returns a claim and a acess_token string
func PostgresQlGenerateAccessClaims(uuid string) (*SignedDetails, string) {

	t := time.Now()
	claim := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(15 * time.Minute).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return claim, tokenString
}

func PostgresQlGenerateRefreshClaims(sd *SignedDetails) string {
	result := database.PostgresDB.Where(&SignedDetails{
		StandardClaims: jwt.StandardClaims{
			Issuer: sd.Issuer,
		},
	}).Find(&SignedDetails{})

	//chceking the number of refresh token stored
	//if the number is higher than 3, remove all the refresh token and leave only one
	if result.RowsAffected > 3 {
		database.PostgresDB.Where(&SignedDetails{
			StandardClaims: jwt.StandardClaims{Issuer: sd.Issuer},
		}).Delete(&SignedDetails{})
	}

	t := time.Now()
	refreshClaim := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			Issuer:    sd.Issuer,
			ExpiresAt: t.Add(30 * 24 * time.Hour).Unix(),
			Subject:   "refresh_token",
			IssuedAt:  t.Unix(),
		},
	}

	//create a claim in the database
	database.PostgresDB.Create(&refreshClaim)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
	}

	return refreshTokenString
}
