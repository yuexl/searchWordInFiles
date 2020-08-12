package rpc

import (
	"context"
	"net/url"
	"time"

	"fileSearch/log"
	"fileSearch/proto"

	"fileSearch/fileSearchRpc/logic"
)

type FileRpcSearch struct {
}

func (rpc *FileRpcSearch) Search(ctx context.Context, req *proto.SearchWordReq, rsp *proto.SearchWordRsp) error {
	word, err := url.QueryUnescape(req.Word)
	if err != nil {
		return err
	}
	log.GLogger.WithField("traceid", req.TraceId).Infoln(word)
	rsp.ServerId = req.TraceId + "_" + time.Now().String()
	logic.StartSearch(req.Word, rsp)

	return nil
}
