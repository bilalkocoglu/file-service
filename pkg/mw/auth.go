package mw

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_const "github.com/imminoglobulin/file-service/pkg/const"
	"github.com/imminoglobulin/file-service/pkg/database"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		basicToken := req.Header.Get(_const.AuthorizationHeader)
		if basicToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token must be not null.",
			})
			return
		}
		tokenType, token := parseToken(basicToken)
		if tokenType != "Basic" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token type must be Basic.",
			})
			return
		}
		err := basicAuthInternal(token, c)
		if err != nil {
			return
		}

		c.Next()
	}
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		bearerToken := req.Header.Get(_const.AuthorizationHeader)

		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token must be not null.",
			})
			return
		}

		tokenType, tokenString := parseToken(bearerToken)

		if tokenType != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token type must be Bearer.",
			})
			return
		}

		err := jwtAuthInternal(tokenString, c)
		if err != nil {
			return
		}

		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		req := context.Request

		token := req.Header.Get(_const.AuthorizationHeader)

		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token must be not null.",
			})
			return
		}

		tokenType, tokenString := parseToken(token)

		if tokenType == "Bearer" {
			err := jwtAuthInternal(tokenString, context)
			if err != nil {
				return
			}
		} else if tokenType == "Basic" {
			err := basicAuthInternal(tokenString, context)
			if err != nil {
				return
			}
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Undefined token type.",
			})
			return
		}

		context.Next()
	}
}

func VerifyJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func parseToken(token string) (string, string) {
	parsedToken := strings.Split(token, " ")

	return parsedToken[0], parsedToken[1]
}

func basicAuthInternal(token string, ctx *gin.Context) error {
	decodeToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token decode exception",
		})
		return err
	}
	splitToken := strings.Split(string(decodeToken), ":")
	username := splitToken[0]
	password := splitToken[1]

	var user database.ApplicationUser
	err = database.GetUserByUsername(&user, username)
	if err != nil || user.ID == 0 || user.Password != password {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return err
	}

	ctx.Set(_const.CurrentUser, user)
	return nil
}

func jwtAuthInternal(tokenString string, ctx *gin.Context) error {
	token, err := VerifyJwtToken(tokenString)
	if err != nil {
		log.Error().Err(err).Msg("Crash verify jwt token.")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		log.Error().Msg("Can not create claims map.")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return err
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("user_id claims not found in jwt.")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return err
	}

	var user database.ApplicationUser
	err = database.GetUserById(&user, userId)
	if err != nil || user.ID == 0 {
		log.Error().Err(err).Msg("User not found for jwt user id")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return err
	}

	ctx.Set(_const.CurrentUser, user)
	return nil
}
