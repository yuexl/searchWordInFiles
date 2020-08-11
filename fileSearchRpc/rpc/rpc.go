package rpc

import (
	"context"
	"time"

	"fileSearch/proto"

	"fileSearch/fileSearchRpc/logic"
	"fileSearch/log"
)

type FileRpcSearch struct {
}

func (rpc *FileRpcSearch) Search(ctx context.Context, req *proto.SearchWordReq, rsp *proto.SearchWordRsp) error {
	log.GLogger.WithField("traceid", req.TraceId).Infoln(req.Word)
	rsp.ServerId = req.TraceId + "_" + time.Now().String()
	logic.StartSearch(req.Word, rsp)

	return nil
}
