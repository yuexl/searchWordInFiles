package rpc

import (
	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/client"

	"fileSearch/api/config"
	"fileSearch/api/log"
)

var GXClient client.XClient

func InitRpcClient() client.XClient {
	etcdAddr := config.GConfig.Etcd.Host + ":" + config.GConfig.Etcd.Port
	serviceDiscovery := client.NewEtcdV3Discovery(config.GConfig.Etcd.BasePath, "FileRpcSearch", []string{etcdAddr}, nil)
	rpcClient := client.NewXClient("FileRpcSearch", client.Failover, client.RoundRobin, serviceDiscovery, client.DefaultOption)
	GXClient = rpcClient

	log.GLogger.WithFields(logrus.Fields{
		"etcd": etcdAddr,
		"path": config.GConfig.Etcd.BasePath,
	}).Infoln("init rpc client")

	return rpcClient
}
