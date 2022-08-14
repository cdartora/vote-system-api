package middleware

import (
	"fmt"
	"net/http"

	helper "example.com/vote-system-api/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token not found"})
		return
	}
	fmt.Println("tokenString: " + tokenString)
	token, err := helper.ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		println(claims)
		ctx.Set("user", claims)
		ctx.Next()
	} else {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token invalid"})
	}
}
