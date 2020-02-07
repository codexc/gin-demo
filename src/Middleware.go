package src

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func MustLogin() gin.HandlerFunc{
    return func(ctx *gin.Context) {
        if _, status := ctx.GetQuery("token"); !status {
            ctx.String(http.StatusUnauthorized, "no token")
            ctx.Abort()
        } else {
            ctx.Next()
        }
    }
}