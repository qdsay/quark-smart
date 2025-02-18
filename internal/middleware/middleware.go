package middleware

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-smart/config"
)

// 中间件
func Handle(ctx *builder.Context) error {
	if config.App.Env == "demo" {
		isForbiddenRoute := false
		forbiddenRoutes := []string{
			"/api/admin/admin/store",
			"/api/admin/admin/save",
			"/api/admin/admin/delete",
			"/api/admin/admin/editable",
			"/api/admin/admin/action/delete",
			"/api/admin/admin/action/change-status",
			"/api/admin/menu/store",
			"/api/admin/menu/save",
			"/api/admin/menu/delete",
			"/api/admin/menu/editable",
			"/api/admin/menu/action/delete",
			"/api/admin/menu/action/change-status",
			"/api/admin/account/action/change-account",
		}
		for _, forbiddenRoute := range forbiddenRoutes {
			if ctx.Path() == forbiddenRoute {
				isForbiddenRoute = true
			}
		}
		if isForbiddenRoute {
			return ctx.JSON(200, msg.Error("演示站点禁止了此操作！", ""))
		}
	}

	return ctx.Next()
}
