package handlers

import (
	"net/url"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"

	"fileSearch/log"
	"fileSearch/proto"

	"fileSearch/api/rpc"
)

func IndexHander(ctx *fiber.Ctx) {
	log.GLogger.WithFields(logrus.Fields{
		"clientIp": ctx.IP(),
		"url":      ctx.BaseURL(),
	}).Infoln("index handler")
	ctx.Send(ctx.BaseURL())
}

func EchoHandle(ctx *fiber.Ctx) {
	word, err := url.QueryUnescape(ctx.Params("word"))
	if err != nil {
		return
	}
	log.GLogger.WithFields(logrus.Fields{
		"clientIp": ctx.IP(),
		"url":      ctx.BaseURL(),
		"param":    word,
	}).Infoln("echo handler")

	ctx.Status(200).Send(word)
}

func GetSearchHandle(ctx *fiber.Ctx) {
	word, err := url.QueryUnescape(ctx.Params("word"))
	if err != nil {
		return
	}

	logFields := logrus.Fields{
		"clientIp": ctx.IP(),
		"url":      ctx.BaseURL(),
		"param":    word,
	}
	log.GLogger.WithFields(logFields).Infoln("search handler")

	req := proto.SearchWordReq{}
	req.TraceId = ctx.IP()
	req.Word = word
	rsp := proto.SearchWordRsp{}
	call, err := rpc.GXClient.Go(ctx.Context(), "Search", &req, &rsp, nil)
	if err != nil {
		log.GLogger.WithFields(logFields).Error(err)
		return
	}
	<-call.Done
	ctx.Status(200).JSON(call.Reply)
	//log.GLogger.Infoln(call.Reply)

	//err := rpc.GXClient.Call(ctx.Context(), "Search", &req, &rsp)
	//if err != nil {
	//	ctx.Status(400).Send("rpc error")
	//}
	//ctx.Status(200).Send(&rsp)
}
