package models

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
