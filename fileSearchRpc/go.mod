module fileSearch/fileSearchRpc

go 1.14

require (
	fileSearch/log v0.0.0
	fileSearch/proto v0.0.0
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/garyburd/redigo v1.6.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/sirupsen/logrus v1.6.0
	github.com/smallnest/rpcx v0.0.0-20200729031544-75f1e2894fdb
	gopkg.in/yaml.v2 v2.2.8
)

replace (
	fileSearch/log => ../log
	fileSearch/proto => ../proto
)
