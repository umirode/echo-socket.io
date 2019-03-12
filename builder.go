package golang_echo_socket_io

import "github.com/labstack/echo"

func BuildServer(serverBuilder func(echo.Context) *Wrapper) echo.HandlerFunc {
	return func(c echo.Context) error {
		serverBuilder(c).Server.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
