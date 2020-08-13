package rpc

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"fileSearch/log"
	"fileSearch/proto"

	"fileSearch/fileSearchRpc/logic"
	"fileSearch/fileSearchRpc/redis"
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
	valStr, err := redis.RedigoStringExec("GET", word)
	if err != nil {
		log.GLogger.Errorln("redis get ", word, err)
	}
	if strings.TrimSpace(valStr) != "" {
		log.GLogger.Infof("redis found %s in %s", word, valStr)
		rsp.Found = true
		err := json.Unmarshal([]byte(valStr), &rsp.SearchRes)
		if err != nil {
			log.GLogger.Errorf("redis found %s in %s , unmarshal err %v", word, valStr, err)
		}
		rsp.FileNum = int64(len(rsp.SearchRes))
	} else {
		logic.StartSearch(req.Word, rsp)
	}
	return nil
}
