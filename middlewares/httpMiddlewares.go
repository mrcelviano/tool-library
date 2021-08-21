package middlewares

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/mrcelviano/tool-library/dbrsession"
)

func HTTPDBRSessionPG(db *dbr.Connection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			ctx = dbrsession.NewContext(ctx, db)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
