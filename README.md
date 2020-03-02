# Wrapper for use Echo context with Socket.io.

[![Build Status](https://travis-ci.com/umirode/echo-socket.io.svg?branch=master)](https://travis-ci.com/umirode/echo-socket.io)

## Documentation
[pkg.go.dev](https://pkg.go.dev/github.com/umirode/echo-socket.io)

## Install

Install the package with:

```bash
go get -u github.com/umirode/echo-socket.io
```

Import it with:

```go
import esi "github.com/umirode/echo-socket.io"
```

and use `esi` inside your code.

## Example

```go
package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	esi "github.com/umirode/echo-socket.io"
)

func main() {
	e := echo.New()

	e.Any("/socket.io/", socketIOWrapper())

	e.Logger.Fatal(e.Start(":8080"))
}

func socketIOWrapper() func(context echo.Context) error {
	wrapper, err := esi.NewWrapper(nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	wrapper.OnConnect("", func(context echo.Context, conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		return nil
	})
	wrapper.OnError("", func(context echo.Context, e error) {
		fmt.Println("meet error:", e)
	})
	wrapper.OnDisconnect("", func(context echo.Context, conn socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	wrapper.OnEvent("", "test", func(context echo.Context, conn socketio.Conn, msg string) {
		conn.SetContext(msg)
		fmt.Println("notice:", msg)
		conn.Emit("test", msg)
	})

	return wrapper.HandlerFunc
}
```

