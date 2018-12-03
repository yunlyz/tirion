package collect

import (
	"github.com/yunlyz/tirion/models"
	"time"
)

type Getter func() ([]*models.IP)

var (
	getters []Getter
)

func Run() {
	var getters = []Getter{
		data5u,
		xici,
		iphai,
	}

	var defaultIPChan = make(chan *models.IP, 1000)
	ticker := time.NewTicker(time.Hour * 1)
	for {
		select {
		case <-ticker.C:
			for _, getter := range getters {
				go func() {
					ips := getter()
					for _, ip := range ips {
						defaultIPChan <- ip
					}
				}()
			}
		default:
		}
	}
}
