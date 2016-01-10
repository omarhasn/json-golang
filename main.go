package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// auto generated json-struct using  http://json2struct.mervine.net
type MyjsonData struct {
	Alt        string `json:"alt"`
	Day        string `json:"day"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      string `json:"month"`
	News       string `json:"news"`
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
}

func main() {

	data := MyjsonData{}

	resp, _ := http.Get("https://xkcd.com/571/info.0.json")
	bs, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal([]byte(bs), &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}
