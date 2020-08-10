package rpc

import (
	"github.com/smallnest/rpcx/client"

	"fileSearch/api/config"
)

var GXClient client.XClient

func InitRpcClient() client.XClient {
	etcdAddr := config.GConfig.Etcd.Host + ":" + config.GConfig.Etcd.Port
	serviceDiscovery := client.NewEtcdV3Discovery(config.GConfig.Etcd.BasePath, "FileRpcSearch", []string{etcdAddr}, nil)
	rpcClient := client.NewXClient("FileRpcSearch", client.Failover, client.RoundRobin, serviceDiscovery, client.DefaultOption)
	GXClient = rpcClient
	return rpcClient
}
