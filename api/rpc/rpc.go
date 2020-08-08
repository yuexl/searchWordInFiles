package rpc

import "github.com/smallnest/rpcx/client"

var GXClient client.XClient

func InitRpcClient() client.XClient {
	serviceDiscovery := client.NewPeer2PeerDiscovery("tcp@:9001", "")
	rpcClient := client.NewXClient("file_search_rpc", client.Failover, client.RoundRobin, serviceDiscovery, client.DefaultOption)
	GXClient = rpcClient
	return rpcClient
}
