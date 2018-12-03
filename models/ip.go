package models

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

const (
	IpProxy = "ip:proxy"
)

type IP struct {
	ID           int64   `bson:"id"`
	Address      string  `bson:"address"`
	Port         int32   `bson:"port"`
	ProtocolType string  `bson:"protocol_type"`
	Delay        float32 `bson:"delay"`
}

func (ip *IP) Insert() (int64, error) {
	client := redis.NewClient(&redis.Options{Addr: "127.0.01:6379", Password: "", DB: 0})
	str, _ := json.Marshal(ip)
	return client.LPush(IpProxy, str).Result()
}

func (ip *IP) Update() (int64, error) {
	client := redis.NewClient(&redis.Options{Addr: "127.0.01:6379", Password: "", DB: 0})
	str, _ := json.Marshal(ip)
	return client.LPush(IpProxy, str).Result()
}

func (ip *IP) Get(index int64) (*IP, error) {
	client := redis.NewClient(&redis.Options{Addr: "127.0.01:6379", Password: "", DB: 0})
	str, err := client.LIndex(IpProxy, index).Result()
	res := &IP{}
	json.Unmarshal([]byte(str), res)
	return res, err

}
