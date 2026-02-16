package public

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	delivery *Delivery
}

func (routes Route) Register(r *gin.Engine) {
	r.GET(RoutePing, routes.delivery.HealthCheck)
}

func NewRoute(delivery *Delivery) *Route {
	return &Route{delivery: delivery}
}
