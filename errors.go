package websocket

import "github.com/aivoicesystems/aivoice/common/errors"

var Error = errors.Make("websocket", nil)

var CloseError = Error.Make("close")

var ClosedError = Error.Make("closed")

var ConnectError = Error.Make("close")

var ReadError = Error.Make("read")

var UpgradeError = Error.Make("upgrade")

var WriteError = Error.Make("write")
