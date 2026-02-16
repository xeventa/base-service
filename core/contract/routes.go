package contract

import "github.com/gin-gonic/gin"

type IRoute interface {
	Register(r *gin.Engine)
}
