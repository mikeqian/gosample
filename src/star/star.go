package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	cookie = "__gads=ID=c8c647ad97702bcb:T=1429676504:S=ALNI_MaZFlwOJP3wKUpFlNbognwO9654QQ; pgv_pvi=8675348480; lzstat_uv=1503350266744680742|760861@2235193; __utma=226521935.2089370079.1429676489.1441180461.1441518078.41; __utmz=226521935.1441518078.41.25.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; .CNBlogsCookie=B933D1A3BDA0B6A5397B1175C6FF3076D4A302E6DD9FC5B315218A61546D33EB53E61AE225BC5830AA128976935A02109266C7280B5B1306A0C9AA3016310597E36D19553523E299CAB3CF1C0E7D360058E16BBE; _gat=1; td_cookie=224553786; _ga=GA1.2.2089370079.1429676489; SERVERID=73ea7682c79ff5c414f1e6047449c5c1|1441960935|1441957193"
)

type Ing struct {
	Content    string
	PublicFlag int32
}

func getLastIng() (text string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ing.cnblogs.com/ajax/ing/GetIngList?IngListType=my&PageIndex=1&PageSize=1&Tag=&_=1441948524646", nil)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}

func insertIng(i int) {
	ing := Ing{}
	ing.PublicFlag = 1
	ing.Content = "mm" + strconv.Itoa(i)
	text, _ := json.Marshal(ing)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://ing.cnblogs.com/ajax/ing/Publish", bytes.NewReader(text))
	if err != nil {
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func deleteIng(ing string) {
	begin := strings.Index(ing, "DelIng(")
	l := len("DelIng(")
	ingId := ing[begin+l : begin+l+6]

	fmt.Println("{ingId:" + ingId + "}")

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://ing.cnblogs.com/ajax/ing/del", strings.NewReader("{ingId:"+ingId+"}"))

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func main() {
	for i := 1; i < 20; i++ {
		ing := getLastIng()
		if strings.Contains(ing, "幸运闪") {
			insertIng(i)
		} else {
			deleteIng(ing)
		}
		time.Sleep(15 * time.Minute)
	}
}
