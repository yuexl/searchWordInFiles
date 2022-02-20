package redis

import (
	"fmt"
	"time"

	redigo "github.com/garyburd/redigo/redis"

	"rpc/config"
)

type RedigoPool struct {
	pool *redigo.Pool
}

var redigoPool *RedigoPool

func InitRedigo() {
	addr := fmt.Sprintf("%s:%s", config.GConfig.Redis.Host, config.GConfig.Redis.Port)
	redigoPool = new(RedigoPool)
	redigoPool.pool = &redigo.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial(
				"tcp",
				addr,
				redigo.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redigo.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redigo.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redigo.DialDatabase(0),
			)
		},
	}
}

func RedigoExec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redigoPool.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

func RedigoStringExec(cmd string, key interface{}, args ...interface{}) (string, error) {
	con := redigoPool.pool.Get()
	if err := con.Err(); err != nil {
		return "", err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return redigo.String(con.Do(cmd, parmas...))
}
