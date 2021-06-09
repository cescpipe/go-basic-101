package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"runtime"
)

func BodyDumpConfig() middleware.BodyDumpConfig {
	handler := func(c echo.Context, req []byte, res []byte) {
		//ctx := c.Request().Context()

		log.Infof("headers: %+v", c.Request().Header)
		log.Infof("request: %s", req)
		log.Infof("response: %s", res)
	}
	return middleware.BodyDumpConfig{Handler: handler}
}

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", rec)
				}
				stack := make([]byte, 4<<10) // 4KB
				length := runtime.Stack(stack, false)

				log.Errorf("[PANIC RECOVER] %v: %s", err, stack[:length])
			}
		}()

		return next(c)
	}
}
