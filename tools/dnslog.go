package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tangx/goutils/md5sha"
)

const (
	DNSLOG = `http://dnslog.cn`
)

type RecordsResp struct {
	Records [][]string `json:"records,omitempty"`
}

func GetDomain() (domain string, cookie string) {
	GetDomainURL := fmt.Sprintf("%s/getdomain.php", DNSLOG)

	session := fmt.Sprintf("PHPSESSID=%s", md5sha.Time())

	resp, err := dnslogGet(GetDomainURL, session)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("resp Error:%s", err.Error())
	}

	domain = fmt.Sprintf("%s", data)
	cookie = resp.Request.Header.Get("Cookie")

	return
}

func GetRecords(session string) RecordsResp {

	GetRecordsURL := fmt.Sprintf("%s/getrecords.php", DNSLOG)

	resp, err := dnslogGet(GetRecordsURL, session)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data)) // 这里不满足 json.Unmarshal 需求
	data = trans(data)

	var records = RecordsResp{}
	err = json.Unmarshal(data, &records)
	if err != nil {
		panic(err)
	}

	return records
}

func (r RecordsResp) IsEmpty() bool {
	// 假设结构体中有一个切片字段 Data
	if len(r.Records) == 0 {
		return true
	} else {
		return false
	}
}

func dnslogGet(url string, session string) (resp *http.Response, err error) {

	//
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", session)

	return http.DefaultClient.Do(req)

}

func trans(body []byte) []byte {
	header := `{"records":`
	footer := `}`

	r := fmt.Sprintf("%s%s%s", header, string(body), footer)

	return []byte(r)
}
