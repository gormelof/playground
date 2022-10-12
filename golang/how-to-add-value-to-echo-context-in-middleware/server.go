package playground

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	App *echo.Echo
}

func (s *Server) NewServer(e *echo.Echo) *Server {
	s.initMiddlewares(e)
	s.initRoutes(e)
	return &Server{App: e}
}

func (s *Server) initMiddlewares(e *echo.Echo) {
	e.Use(s.setCtxValuesMiddleware())
}

func (s *Server) initRoutes(e *echo.Echo) {
	e.GET("/users/:id", s.getUserById)
}

func (s *Server) getUserById(c echo.Context) error {
	fmt.Printf("correlation-id from request context => %s /n", c.Request().Context().Value("correlation-id"))
	return c.String(http.StatusOK, "")
}

func (s *Server) setCtxValuesMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "correlation-id", uuid.New().String())
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
