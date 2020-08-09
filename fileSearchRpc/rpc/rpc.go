package rpc

import (
	"context"
	"fmt"
	"time"

	"fileSearch/fileSearchRpc/logic"
	"fileSearch/fileSearchRpc/proto"
)

type FileRpcSearch struct {
}

func (rpc *FileRpcSearch) Search(ctx context.Context, req *proto.SearchWordReq, rsp *proto.SearchWordRsp) error {
	fmt.Println(req.TraceId, req.Word)
	rsp.ServerId = req.TraceId + "_" + time.Now().String()
	logic.StartSearch(req.Word, rsp)

	return nil
}
