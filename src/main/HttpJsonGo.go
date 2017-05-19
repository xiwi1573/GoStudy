package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func main() {
	//http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json   根据IP获取城市信息   https://xkcd.com/571/info.0.json
	// url := "https://cdn.heweather.com/china-city-list.json"
	url := "https://xkcd.com/571/info.0.json"

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", string(body))

	js, err := simplejson.NewJson(body)
	if err != nil {
		log.Fatalln(err)
	}

	num := js.Get("num").MustInt()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("num:%d", num)
}
