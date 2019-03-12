package golang_echo_socket_io

import (
	"errors"
	"github.com/googollee/go-engine.io"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

/**
Socket.io server interface
*/
type ISocketIO interface {
	http.Handler
	io.Closer

	Serve() error

	OnConnect(nsp string, f func(socketio.Conn) error)
	OnDisconnect(nsp string, f func(socketio.Conn, string))
	OnError(nsp string, f func(error))
	OnEvent(nsp, event string, f interface{})
}

/**
Socket.io wrapper interface
*/
type ISocketIOWrapper interface {
	OnConnect(context echo.Context, nsp string, f func(echo.Context, socketio.Conn) error)
	OnDisconnect(context echo.Context, nsp string, f func(echo.Context, socketio.Conn, string))
	OnError(context echo.Context, nsp string, f func(echo.Context, error))
	OnEvent(context echo.Context, nsp, event string, f func(echo.Context, socketio.Conn, string))
}

type Wrapper struct {
	Server ISocketIO
}

/**
Create wrapper and Socket.io server
*/
func NewWrapper(options *engineio.Options) (*Wrapper, error) {
	server, err := socketio.NewServer(options)
	if err != nil {
		return nil, err
	}

	return &Wrapper{
		Server: server,
	}, nil
}

/**
Create wrapper with exists Socket.io server
*/
func NewWrapperWithServer(server ISocketIO) (*Wrapper, error) {
	if server == nil {
		return nil, errors.New("socket.io server can not be nil")
	}

	return &Wrapper{
		Server: server,
	}, nil
}

/**
On Socket.io client connect
*/
func (s *Wrapper) OnConnect(context echo.Context, nsp string, f func(echo.Context, socketio.Conn) error) {
	s.Server.OnConnect(nsp, func(conn socketio.Conn) error {
		return f(context, conn)
	})
}

/**
On Socket.io client disconnect
*/
func (s *Wrapper) OnDisconnect(context echo.Context, nsp string, f func(echo.Context, socketio.Conn, string)) {
	s.Server.OnDisconnect(nsp, func(conn socketio.Conn, msg string) {
		f(context, conn, msg)
	})
}

/**
On Socket.io error
*/
func (s *Wrapper) OnError(context echo.Context, nsp string, f func(echo.Context, error)) {
	s.Server.OnError(nsp, func(e error) {
		f(context, e)
	})
}

/**
On Socket.io event from client
*/
func (s *Wrapper) OnEvent(context echo.Context, nsp, event string, f func(echo.Context, socketio.Conn, string)) {
	s.Server.OnEvent(nsp, event, func(conn socketio.Conn, msg string) {
		f(context, conn, msg)
	})
}
