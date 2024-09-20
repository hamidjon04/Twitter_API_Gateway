package middleware

import (
	"api/api/token"
	"api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, model.Error{
				Message: "authorization header is required",})
			c.Abort()
			return
		}

		valid, err := token.ValidToken(auth)
		if !valid || err != nil {
			c.JSON(http.StatusUnauthorized, model.Error{Message: fmt.Sprintf("invalid token: %v", err)})
			c.Abort()
			return
		}

		claims, err := token.ExtractClaimToken(auth)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, model.Error{Message: fmt.Sprintf("invalid token claims: %v", err)})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

// func Authorize(enforcer *casbin.Enforcer) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		access := ctx.GetHeader("Authorization")
// 		claim, err := token.ExtractClaimToken(access)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
// 			return
// 		}
// 		log.Println(claim.Role, ctx.FullPath(), ctx.Request.Method)
// 		ok, err := enforcer.Enforce(claim.Role, ctx.FullPath(), ctx.Request.Method)
// 		if err != nil {
// 			fmt.Println(err)
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"Error": "Internal server error",
// 				"Err":   err.Error(),
// 			})
// 			ctx.Abort()
// 			return
// 		}

// 		if !ok {
// 			ctx.JSON(http.StatusForbidden, gin.H{
// 				"Error": "Forbidden",
// 			})
// 			ctx.Abort()
// 			return
// 		}

// 		ctx.Next()

// 	}
// }
