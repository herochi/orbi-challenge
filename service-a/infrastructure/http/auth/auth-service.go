package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github/herochi/orbi/service-a/domain/auth"
)

type AuthService interface {
	ValidatePassword(password string, enteredPassword string) error
	GenerateToken(userData *auth.UserData) *auth.TokenDetail
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	UserData *auth.UserData `json:"userdata"`
	jwt.StandardClaims
}

type authServices struct {
	secretKey        string
	refreshSecretKey string
	issure           string
}

func NewAuthService() AuthService {
	return &authServices{
		secretKey:        getSecretKey("access"),
		refreshSecretKey: getSecretKey("refresh"),
		issure:           "gst-report-api",
	}
}

func getSecretKey(secreteType string) string {
	key := "JWT_SECRET"

	if secreteType == "refresh" {
		key = "REFRESH_JWT_SECRET"
	}

	secret := viper.GetString(key)

	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *authServices) ValidatePassword(password string, enteredPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(enteredPassword)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	return nil
}

func (service *authServices) GenerateToken(userData *auth.UserData) *auth.TokenDetail {
	td := &auth.TokenDetail{}

	var err error

	atClaims := &authCustomClaims{
		UserData: userData,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    service.issure,
			NotBefore: 0,
			Subject:   "",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = accessToken.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}

	rtClaims := &authCustomClaims{
		UserData: userData,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    service.issure,
			NotBefore: 0,
			Subject:   "",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = refreshToken.SignedString([]byte(service.refreshSecretKey))
	if err != nil {
		panic(err)
	}

	return td
}

func (service *authServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
