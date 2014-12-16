package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	body := `{"title":"afe1112311Ff","private":true}`
	fmt.Println(body)
	body_reader := strings.NewReader(body)
	req, _ := http.NewRequest("POST", "http://api.heclouds.com/devices", body_reader)
	// ...
	req.Header.Add("api-key", "gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	// ...
	client := &http.Client{}
	resp, _ := client.Do(req)
	//	client.Do(req)
	fmt.Println(resp.ContentLength)
	b := make([]byte, resp.ContentLength)

	resp.Body.Read(b)
	fmt.Println(string(b))

	var r interface{}
	err := json.Unmarshal(b, &r)
	if err == nil {
		g, ok := r.(map[string]interface{})
		if ok {
			for k, v := range g {
				switch v2 := v.(type) {
				case string:
					fmt.Println(k, "is ", v2)
				case float64:
					fmt.Println(k, "is ", v2)
				case bool:
					fmt.Println(k, "is ", v2)
				case []interface{}:
					fmt.Println(k, "is ", v2)
				case map[string]interface{}:
					fmt.Println(k, "is ", v2)
					xx, _ := json.Marshal(v2)
					fmt.Println (string(xx))
				default:
					fmt.Println(k, "is other")
				}
			}
		}

	}
}
