module github.com/ab36245/go-websocket

go 1.24.2

require (
	github.com/ab36245/go-errors v0.0.1
	github.com/ab36245/go-stream v0.0.1
	github.com/ab36245/go-writer v0.0.1
	github.com/gorilla/websocket v1.5.3
)

replace github.com/ab36245/go-errors => ../go-errors

replace github.com/ab36245/go-stream => ../go-stream

replace github.com/ab36245/go-writer => ../go-writer
