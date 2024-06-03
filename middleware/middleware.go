package middleware

import (
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error(), "Unauthenticated"))
			return
		}
		c.Set("userInfo", verifyToken)
		c.Next()
	}
}