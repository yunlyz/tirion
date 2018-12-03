package collect

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/yunlyz/tirion/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko)" +
		" Chrome/71.0.3578.62 Safari/537.36"
)

func data5u() []*models.IP {
	var url = "http://www.data5u.com/free/gngn/index.shtml"
	var ips []*models.IP

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("User-Agent", UserAgent)
	request.Header.Set("Referer", url)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find(".wlist > ul > li").Next().Find("ul").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			ip := &models.IP{}
			selection.Find("span").Each(func(i int, span *goquery.Selection) {
				switch i {
				case 0:
					ip.Address = span.Text()
				case 1:
					port, _ := strconv.Atoi(span.Text())
					ip.Port = int32(port)
				case 3:
					ip.ProtocolType = span.Text()
				case 7:
					str := strings.Split(span.Text(), " ")
					delay, _ := strconv.ParseFloat(str[0], 32)
					ip.Delay = float32(delay)
				default:
				}
			})
			ips = append(ips, ip)
		}
	})
	return ips
}

func xici() []*models.IP {
	var url = "http://www.xicidaili.com/nn/1"
	var ips []*models.IP

	request, _ := http.NewRequest(http.MethodGet, url, nil)

	request.Header.Set("Host", "www.xicidaili.com")
	request.Header.Set("User-Agent", UserAgent)
	request.Header.Set("Referer", "http://www.xicidaili.com/")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find("#ip_list > tbody > tr").Each(func(i int, tr *goquery.Selection) {
		if i != 0 {
			ip := &models.IP{}
			tr.Find("td").Each(func(i int, selection *goquery.Selection) {
				switch i {
				case 1:
					ip.Address = selection.Text()
				case 2:
					port, _ := strconv.Atoi(selection.Text())
					ip.Port = int32(port)
				case 5:
					ip.ProtocolType = selection.Text()
				case 6:
					val, _ := selection.Find("div").Attr("title")
					str := strings.TrimRight(val, "秒")
					delay, _ := strconv.ParseFloat(str, 32)
					ip.Delay = float32(delay)
				default:
				}
			})
			ips = append(ips, ip)
		}
	})
	return ips
}

func iphai() []*models.IP {
	var url = "http://www.iphai.com/free/ng"
	var ips []*models.IP

	request, _ := http.NewRequest(http.MethodGet, url, nil)

	request.Header.Set("Host", "www.xicidaili.com")
	request.Header.Set("User-Agent", UserAgent)
	request.Header.Set("Referer", "http://www.xicidaili.com/")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find("table > tbody > tr").Each(func(i int, tr *goquery.Selection) {
		if i != 0 {
			ip := &models.IP{}
			tr.Find("td").Each(func(i int, selection *goquery.Selection) {
				switch i {
				case 0:
					ip.Address = strings.TrimSpace(selection.Text())
				case 1:
					port, _ := strconv.Atoi(strings.TrimSpace(selection.Text()))
					ip.Port = int32(port)
				case 3:
					ip.ProtocolType = strings.TrimSpace(selection.Text())
				case 5:
					val := selection.Text()
					str := strings.TrimRight(strings.TrimSpace(val), "s")
					delay, _ := strconv.ParseFloat(str, 32)
					ip.Delay = float32(delay)
				default:
				}
			})
			ips = append(ips, ip)
		}
	})
	return ips
}
