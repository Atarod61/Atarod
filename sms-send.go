package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://panel.asanak.com/webservice/v1rest/msgstatus"
	str := "username=YOUR_USER&password=YOUR_PASSWORD&msgid=YOUR_MESSAGE_ID"
	payload := strings.NewReader(str)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
