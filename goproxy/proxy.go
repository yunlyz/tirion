package goproxy

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/lunny/log"
	"github.com/yunlyz/tirion/models"
	"math/rand"
	"net/http"
)

const (
	DefaultCount = 5
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Data struct {
	List       []interface{} `json:"list"`
	Pagination struct {
		CurrentPage int `json:"current_page"`
		PageSize    int `json:"page_size"`
		TotalCount  int `json:"total_count"`
	} `json:"pagination"`
}

func Start() {
	http.HandleFunc("/v1/proxy", handleProxy)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleProxy(writer http.ResponseWriter, request *http.Request) {
	var body []byte
	var ips []*models.IP

	switch request.Method {
	case http.MethodGet:
		client := redis.NewClient(&redis.Options{Addr: "127.0.01:6379", Password: "", DB: 0})
		cnt, _ := client.LLen(models.IpProxy).Result()
		if cnt != 0 {
			for i := 0; i < DefaultCount; i++ {
				str, err := client.LIndex(models.IpProxy, rand.Int63n(cnt)).Result()
				if err != nil {
					log.Error(err)
					continue
				}
				ip := &models.IP{}
				json.Unmarshal([]byte(str), ip)
				ips = append(ips, ip)
			}
		}
		body, _ = json.Marshal(&Result{Code: 0, Message: "success", Data: ips})
	default:
		body, _ = json.Marshal(&Result{Code: 50000, Message: "未支持除GET方法外的其他HTTP方法"})
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
}
