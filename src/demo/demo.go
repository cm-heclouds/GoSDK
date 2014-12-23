package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	//	"net/http"
	//	"strings"
)

func main() {
	a := map[string]interface{}{"title": "afe1112311Ff", "private": true}

	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))

	b := map[string]map[string]interface{}{"1": {"title": "afe1112311Ff", "private": true}, "2": {"title": "afe1112311Ff", "private": true}}
	bytes2, _ := json.Marshal(b)
	fmt.Println(string(bytes2))
	
	
	data_map := make(map[string]interface{})
	data_map["title"] = "title"
	
	permissions:= make([]map[string]interface{}, 1)
	res := make(map[string]interface{})
	
	r:= make([]map[string]interface{}, 1)
	inner_r:=make(map[string]interface{})
	inner_r["dev_id"] = 424
	inner_r["ds_id"] = "fan1"
	r[0] = inner_r
	
	res["resources"] = r
	permissions[0] =res
	data_map["permissions"] = permissions
	bytes3, _ := json.Marshal(data_map)
	fmt.Println(string(bytes3))
	
	//
	//	body_reader := strings.NewReader(string(bytes))
	//	req, _ := http.NewRequest("POST", "http://api.heclouds.com/devices", body_reader)
	//	// ...
	//	req.Header.Add("api-key", "gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	//	// ...
	//	client := &http.Client{}
	//	resp, _ := client.Do(req)
	//	//	client.Do(req)
	//	fmt.Println(resp.ContentLength)
	//	b := make([]byte, resp.ContentLength)
	//
	//	resp.Body.Read(b)
	//	fmt.Println(string(b))
	//
	//	var r interface{}
	//	err := json.Unmarshal(b, &r)
	//	if err == nil {
	//		g, ok := r.(map[string]interface{})
	//		if ok {
	//			for k, v := range g {
	//				switch v2 := v.(type) {
	//				case string:
	//					fmt.Println(k, "is ", v2)
	//				case float64:
	//					fmt.Println(k, "is ", v2)
	//				case bool:
	//					fmt.Println(k, "is ", v2)
	//				case []interface{}:
	//					fmt.Println(k, "is ", v2)
	//				case map[string]interface{}:
	//					fmt.Println(k, "is ", v2)
	//					xx, _ := json.Marshal(v2)
	//					fmt.Println(string(xx))
	//				default:
	//					fmt.Println(k, "is other")
	//				}
	//			}
	//		}
	//
	//	}
}
