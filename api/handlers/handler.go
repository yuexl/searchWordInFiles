package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/smallnest/rpcx/log"

	"fileSearch/api/proto"
	"fileSearch/api/rpc"
)

func IndexHander(ctx *fiber.Ctx) {

	ctx.Send(ctx.BaseURL())
}

func GetSearchHandle(ctx *fiber.Ctx) {
	word := ctx.Params("word")

	req := proto.SearchWordReq{}
	req.TraceId = ctx.IP()
	req.Word = word
	rsp := proto.SearchWordRsp{}
	call, err := rpc.GXClient.Go(ctx.Context(), "Search", &req, &rsp, nil)
	if err != nil {
		return
	}
	<-call.Done
	ctx.Status(200).JSON(call.Reply)
	log.Info(call.Reply)
	log.Info(call)

	//err := rpc.GXClient.Call(ctx.Context(), "Search", &req, &rsp)
	//if err != nil {
	//	ctx.Status(400).Send("rpc error")
	//}
	//ctx.Status(200).Send(&rsp)
}
