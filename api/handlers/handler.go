package handlers

import (
	"net/url"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"

	"fileSearch/log"
	"fileSearch/proto"

	"fileSearch/api/rpc"
	"fileSearch/api/utils"
)

//IndexHander indexccc
func IndexHander(ctx *fiber.Ctx) {
	log.GLogger.WithFields(logrus.Fields{
		"clientIp": ctx.IP(),
		"url":      ctx.BaseURL(),
	}).Infoln("index handler")
	ctx.Send(ctx.BaseURL())
}

func LoginHandle(ctx *fiber.Ctx) {
	uname := ctx.FormValue("username")
	pwd := ctx.FormValue("passwd")
	strGuid := utils.GetGUID().Hex()
	ctx.Cookie(&fiber.Cookie{Name: "sessionid", Value: strGuid})
	ctx.Status(200).Send(uname + " " + pwd + "   login succ")
}

func SessionHandle(ctx *fiber.Ctx) {
	sessionid := ctx.Cookies("sessionid", "")
	print(sessionid)
}

//EchoHandle indexccc
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

//SayHandle indexccc
func SayHandle(ctx *fiber.Ctx) {
	query := ctx.Query("word", "")
	log.GLogger.Info("query ", query)
	word, err := url.QueryUnescape(query)
	if err != nil {
		return
	}
	log.GLogger.WithFields(logrus.Fields{
		"clientIp": ctx.IP(),
		"url":      ctx.BaseURL(),
		"param":    word,
	}).Infoln("say handler")

	ctx.Status(200).Send("hello " + word)
}

func GetSearchHandle(ctx *fiber.Ctx) {
	word, err := url.QueryUnescape(ctx.Query("word"))
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
