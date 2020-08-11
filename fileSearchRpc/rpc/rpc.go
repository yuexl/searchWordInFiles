package rpc

import (
	"context"
	"time"

	"fileSearch/log"
	"fileSearch/proto"

	"fileSearch/fileSearchRpc/logic"
)

type FileRpcSearch struct {
}

func (rpc *FileRpcSearch) Search(ctx context.Context, req *proto.SearchWordReq, rsp *proto.SearchWordRsp) error {
	log.GLogger.WithField("traceid", req.TraceId).Infoln(req.Word)
	rsp.ServerId = req.TraceId + "_" + time.Now().String()
	logic.StartSearch(req.Word, rsp)

	return nil
}
