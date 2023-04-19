package commands

import (
	"strings"

	"github.com/IbrahimMohammed47/codecrafters-redis-go/resp"
)

func HandleCommand(respArr resp.Resp) resp.Resp {
	arr := respArr.(*resp.Array)
	if len(arr.Value) == 0 {
		return resp.NewErrorWithString("first argument must be a command")
	}
	cmdToken := arr.Value[0]
	if cmdToken.Type().String() != "<bulkbytes>" {
		return resp.NewErrorWithString("first argument must be a command")
	} else {
		var res resp.Resp
		cmd := strings.ToUpper(string(cmdToken.(*resp.BulkBytes).Value))
		switch cmd {
		default:
			res = resp.NewErrorWithString("first argument must be a command")
		case "ECHO":
			res = EchoCommand(arr.Value[1:])
		case "PING":
			res = PingCommand()
		}
		return res
	}
}

func EchoCommand(respArr []resp.Resp) resp.Resp {
	return respArr[0]
}

func PingCommand() resp.Resp {
	return resp.NewString("PONG")
}
