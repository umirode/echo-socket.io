package golang_echo_socket_io

import "github.com/labstack/echo"

type SocketIOServerBuilder func(echo.Context) *Wrapper

func BuildServer(serverBuilder SocketIOServerBuilder) echo.HandlerFunc {
	return func(c echo.Context) error {
		serverBuilder(c).Server.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
