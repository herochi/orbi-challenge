package middleware

import (
	"errors"
	"net/http"
	"strings"

	//"github/herochi/orbi/service-a/domain/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type route struct {
	path   string
	method string
}

var routes = []route{
	{"/api/v1/login/", http.MethodPost},
	{"/api/v1/users/", http.MethodPost},
}

var JWT gin.HandlerFunc = func(c *gin.Context) {

	if ok := isAuthRoute(routes, &route{c.FullPath(), c.Request.Method}); !ok {
		return
	}

	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Missing auth token"})
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid/Malformed auth token"})
		return
	}
	if headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid auth token"})
		return
	}

	//tokenString := headerParts[1]
	//claims := &auth.UserClaims{}
	/*token, err := verifyToken(tokenString, claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid auth token"})
		return
	}

	c.Set(viper.GetString("CLAIM.USER"), claims.UserData)*/
}

func verifyToken(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, errors.New("Malformed auth token")
	}
	return token, nil
}

func isAuthRoute(routes []route, route *route) bool {
	for _, r := range routes {
		if r == *route {
			return false
		}
	}
	return true
}
