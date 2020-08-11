package rpc

import (
	"context"
	"time"

	"fileSearch/fileSearchRpc/log"
	"fileSearch/fileSearchRpc/logic"
	"fileSearch/fileSearchRpc/proto"
)

type FileRpcSearch struct {
}

func (rpc *FileRpcSearch) Search(ctx context.Context, req *proto.SearchWordReq, rsp *proto.SearchWordRsp) error {
	log.GLogger.WithField("traceid", req.TraceId).Infoln(req.Word)
	rsp.ServerId = req.TraceId + "_" + time.Now().String()
	logic.StartSearch(req.Word, rsp)

	return nil
}
