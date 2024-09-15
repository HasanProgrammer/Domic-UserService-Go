package WebAPIMiddleware

import (
	"Domic.Infrastructure/Concretes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")

		jwtToken := InfrastructureConcrete.NewJsonWebToken()

		jwtTokenVerification := jwtToken.Verify(token)

		if token != "" && jwtTokenVerification.Error != nil && jwtTokenVerification.Result == true {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, jwtTokenVerification.Error.Error())
	}
}
