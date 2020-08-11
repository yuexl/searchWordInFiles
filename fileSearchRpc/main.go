package main

import (
	"fmt"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"

	"fileSearch/fileSearchRpc/config"
	"fileSearch/fileSearchRpc/rpc"
	"fileSearch/log"
)

func main() {
	rpcSer := server.NewServer()

	addRegistryPlugin(rpcSer)
	rpcSer.RegisterName("FileRpcSearch", new(rpc.FileRpcSearch), "")

	addr := fmt.Sprintf("%s:%s", config.GConfig.Rpc.Host, config.GConfig.Rpc.Port)

	log.GLogger.WithField("addr", addr).Infoln("start rpc server")
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
		log.GLogger.Errorln(err)
	}
	log.GLogger.WithFields(logrus.Fields{
		"addretcd":   r.EtcdServers,
		"addrserver": r.ServiceAddress,
		"basepath":   r.BasePath,
	}).Infoln("add etcd discovery plugin")
	s.Plugins.Add(r)
}
