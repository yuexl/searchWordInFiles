package main

import (
	"github.com/smallnest/rpcx/server"

	"fileSearch/fileSearchRpc/rpc"
)

func main() {
	rpcSer := server.NewServer()
	rpcSer.RegisterName("file_search_rpc", new(rpc.FileRpcSearch), "")

	rpcSer.Serve("tcp", ":9001")
}
