package golang_echo_socket_io

import (
	"errors"
	"github.com/googollee/go-engine.io"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

type SocketIOOnConnectHandler func(echo.Context, socketio.Conn) error
type SocketIOOnDisconnectHandler func(echo.Context, socketio.Conn, string)
type SocketIOOnErrorHandler func(echo.Context, error)
type SocketIOOnEventHandler func(echo.Context, socketio.Conn, string)

type Wrapper struct {
	Server *socketio.Server
}

func NewWrapper(options *engineio.Options) (*Wrapper, error) {
	server, err := socketio.NewServer(options)
	if err != nil {
		return nil, err
	}

	return &Wrapper{
		Server: server,
	}, nil
}

func NewWrapperWithServer(server *socketio.Server) (*Wrapper, error) {
	if server == nil {
		return nil, errors.New("socket.io server can not be nil")
	}

	return &Wrapper{
		Server: server,
	}, nil
}

func (s *Wrapper) OnConnect(context echo.Context, nsp string, f SocketIOOnConnectHandler) {
	s.Server.OnConnect(nsp, func(conn socketio.Conn) error {
		return f(context, conn)
	})
}

func (s *Wrapper) OnDisconnect(context echo.Context, nsp string, f SocketIOOnDisconnectHandler) {
	s.Server.OnDisconnect(nsp, func(conn socketio.Conn, msg string) {
		f(context, conn, msg)
	})
}

func (s *Wrapper) OnError(context echo.Context, nsp string, f SocketIOOnErrorHandler) {
	s.Server.OnError(nsp, func(e error) {
		f(context, e)
	})
}

func (s *Wrapper) OnEvent(context echo.Context, nsp, event string, f SocketIOOnEventHandler) {
	s.Server.OnEvent(nsp, event, func(conn socketio.Conn, msg string) {
		f(context, conn, msg)
	})
}
