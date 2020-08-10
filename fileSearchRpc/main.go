package main

import (
	"fmt"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"

	"fileSearch/fileSearchRpc/config"
	"fileSearch/fileSearchRpc/rpc"
)

func main() {
	rpcSer := server.NewServer()

	addRegistryPlugin(rpcSer)
	rpcSer.RegisterName("FileRpcSearch", new(rpc.FileRpcSearch), "")

	addr := fmt.Sprintf("%s:%s", config.GConfig.Rpc.Host, config.GConfig.Rpc.Port)
	fmt.Println(addr)
	rpcSer.Serve("tcp", addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + fmt.Sprintf("%s:%s", config.GConfig.Rpc.Host, config.GConfig.Rpc.Port),
		EtcdServers:    []string{config.GConfig.Etcd.Host + ":" + config.GConfig.Etcd.Port},
		BasePath:       config.GConfig.Etcd.BasePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
