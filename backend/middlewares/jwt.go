package middlewares

import (
	"errors"
	"fmt"
	"management-backend/models"
	"management-backend/responses"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTMiddleware struct {
	JWTSecret string
}

var jwtMiddleware JWTMiddleware

func InitJWT(jwtSecret string) {
	jwtMiddleware = JWTMiddleware{
		JWTSecret: jwtSecret,
	}
}

func GetJWT() *JWTMiddleware {
	return &jwtMiddleware
}

func (jm *JWTMiddleware) AuthHandler(authRoles ...models.WorkerRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Token, If WS get Token from URL
		var token string
		if c.Request.Header.Get("Upgrade") == "websocket" {
			token = c.Query("token")
		} else {
			token = c.Request.Header.Get("Authorization")
		}

		// Check if toke in correct format
		// ie Bearer xx03xllasx
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}

		// Validate token
		claims, err := jm.validateToken(t[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}

		// Authorize
		authRolesIds := []string{}
		for _, role := range authRoles {
			authRolesIds = append(authRolesIds, strconv.Itoa(int(role.ID)))
		}

		if !contains(authRolesIds, fmt.Sprintf("%v", claims["role_id"])) {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}

		// Set variables
		c.Set("id", claims["id"])
		c.Set("login", claims["login"])
		c.Set("role_id", claims["role_id"])

		c.Next()
	}
}

func (jm *JWTMiddleware) GenerateToken(login string, roleId uint, id uint) (string, error) {
	// New Token
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["login"] = login

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Set worker role
	claims["role_id"] = roleId

	token.Claims = claims

	// Sign and get as a string
	tokenString, err := token.SignedString([]byte(jm.JWTSecret))

	return tokenString, err
}

func (jm *JWTMiddleware) validateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jm.JWTSecret), nil
	})

	if err != nil || (int64(claims["exp"].(float64)) < time.Now().Unix() && claims["exp"].(float64) != -1) {
		return nil, errors.New("Error validate token")
	}

	return claims, nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
