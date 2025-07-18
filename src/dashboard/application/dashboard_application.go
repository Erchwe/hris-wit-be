package application

import (
	"github.com/labstack/echo/v4"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/dashboard/service" // Pastikan path ini benar
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteDashboard(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewDashboardService(s.GetDB(), cfg)

	d := e.Group("/dashboard")
	d.GET("", getDashboardData(svc))
}

func getDashboardData(svc *service.DashboardService) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := svc.GetDashboardData(c.Request().Context())
		if err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed to get dashboard data")
			return err // Error sudah dibungkus di service
		}
		return httpservice.ResponseData(c, result, nil)
	}
}
