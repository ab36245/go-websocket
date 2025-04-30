module github.com/ab36245/go-websocket

go 1.24.2

require (
	github.com/ab36245/go-errors v0.0.0-20250428061939-8b056c3b905e
	github.com/ab36245/go-stream v0.0.0-20250428060212-0ca3d5e03357
	github.com/ab36245/go-writer v0.0.0-20250428065328-235ec24d23f1
	github.com/gorilla/websocket v1.5.3
)

replace github.com/ab36245/go-errors => ../go-errors

replace github.com/ab36245/go-stream => ../go-stream

replace github.com/ab36245/go-writer => ../go-writer
