package golang_echo_socket_io

import (
	"errors"
	"github.com/googollee/go-engine.io"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

/**
Socket.io wrapper interface
*/
type ISocketIOWrapper interface {
	OnConnect(nsp string, f func(echo.Context, socketio.Conn) error)
	OnDisconnect(nsp string, f func(echo.Context, socketio.Conn, string))
	OnError(nsp string, f func(echo.Context, error))
	OnEvent(nsp, event string, f func(echo.Context, socketio.Conn, string))
}

type Wrapper struct {
	Context echo.Context
	Server  *socketio.Server
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
func NewWrapperWithServer(server *socketio.Server) (*Wrapper, error) {
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
func (s *Wrapper) OnConnect(nsp string, f func(echo.Context, socketio.Conn) error) {
	s.Server.OnConnect(nsp, func(conn socketio.Conn) error {
		return f(s.Context, conn)
	})
}

/**
On Socket.io client disconnect
*/
func (s *Wrapper) OnDisconnect(nsp string, f func(echo.Context, socketio.Conn, string)) {
	s.Server.OnDisconnect(nsp, func(conn socketio.Conn, msg string) {
		f(s.Context, conn, msg)
	})
}

/**
On Socket.io error
*/
func (s *Wrapper) OnError(nsp string, f func(echo.Context, error)) {
	s.Server.OnError(nsp, func(e error) {
		f(s.Context, e)
	})
}

/**
On Socket.io event from client
*/
func (s *Wrapper) OnEvent(nsp, event string, f func(echo.Context, socketio.Conn, string)) {
	s.Server.OnEvent(nsp, event, func(conn socketio.Conn, msg string) {
		f(s.Context, conn, msg)
	})
}

/**
Handler function
*/
func (s *Wrapper) HandlerFunc(context echo.Context) error {
	s.Context = context
	s.Server.ServeHTTP(context.Response(), context.Request())
	return nil
}
