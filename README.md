# golang-echo-socket.io

Wrapper for user Echo context with Socket.io.

## Install

Install the package with:

```bash
go get github.com/umirode/golang-echo-socket.io
```

Import it with:

```go
import "github.com/umirode/golang-echo-socket.io"
```

and use `golang_echo_socket_io` as the package name inside the code.

## Dependencies

* [go-socket.io](https://github.com/googollee/go-socket.io)
* [echo](https://github.com/labstack/echo)

## Example

```go
package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"github.com/umirode/golang-echo-socket.io"
)

func main() {
	e := echo.New()

	e.Any("/socket.io", golang_echo_socket_io.BuildServer(builSocketIoServer))

	e.Logger.Fatal(e.Start(":8080"))
}

func builSocketIoServer(context echo.Context) *golang_echo_socket_io.Wrapper {
	wrapper, err := golang_echo_socket_io.NewWrapper(nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	wrapper.OnConnect(context, "/", func(context echo.Context, conn socketio.Conn) error {
		fmt.Println("connected:", conn.ID())
		return nil
	})
	wrapper.OnError(context, "/", func(context echo.Context, e error) {
		fmt.Println("meet error:", e)
	})
	wrapper.OnDisconnect(context, "/", func(context echo.Context, conn socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	wrapper.OnEvent(context, "/", "test", func(context echo.Context, conn socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		conn.Emit("test", msg)
	})

	go wrapper.Server.Serve()

	return wrapper
}

```

