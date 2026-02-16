package public

import (
	"github.com/gin-gonic/gin"
)

// HealthHandler exposes HTTP endpoints backed by an abstract HealthUseCase.

type Delivery struct {
	service IService
}

func NewDelivery(service IService) *Delivery {
	return &Delivery{service: service}
}

func (d *Delivery) HealthCheck(ctx *gin.Context) {
	data := d.service.HealthCheck()
	ctx.JSON(200, data)
}
