package echo_socket_io_test

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	echo_socket_io "github.com/umirode/echo-socket.io"
	"testing"
)

func TestNewWrapper(t *testing.T) {
	wrapper, err := echo_socket_io.NewWrapper(nil)

	assert.NotNil(t, wrapper)
	assert.Nil(t, err)
}

func TestNewWrapperWithServer(t *testing.T) {
	socketioServer, err := socketio.NewServer(nil)

	assert.NotNil(t, socketioServer)
	assert.Nil(t, err)

	wrapper, err := echo_socket_io.NewWrapperWithServer(socketioServer)

	assert.NotNil(t, wrapper)
	assert.Nil(t, err)
}

func TestWrapper_OnConnect(t *testing.T) {
	wrapper, err := echo_socket_io.NewWrapper(nil)

	assert.NotNil(t, wrapper)
	assert.Nil(t, err)

	wrapper.OnConnect("", func(context echo.Context, conn socketio.Conn) error {
		return nil
	})
}

func TestWrapper_OnDisconnect(t *testing.T) {
	wrapper, err := echo_socket_io.NewWrapper(nil)

	assert.NotNil(t, wrapper)
	assert.Nil(t, err)

	wrapper.OnDisconnect("", func(context echo.Context, conn socketio.Conn, s string) {

	})
}

func TestWrapper_OnError(t *testing.T) {
	wrapper, err := echo_socket_io.NewWrapper(nil)

	assert.NotNil(t, wrapper)
	assert.Nil(t, err)

	wrapper.OnError("", func(context echo.Context, e error) {

	})
}

func TestWrapper_OnEvent(t *testing.T) {
	wrapper, err := echo_socket_io.NewWrapper(nil)

	assert.NotNil(t, wrapper)
	assert.Nil(t, err)

	wrapper.OnEvent("", "", func(echo.Context, socketio.Conn, string) {

	})
}
