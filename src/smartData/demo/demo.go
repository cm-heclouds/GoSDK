package main

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
	"smartData"
	//	"net/http"
	//	"strings"
)

func main() {
	test()
	//	a := map[string]interface{}{"title": "afe1112311Ff", "private": true}
	//
	//	bytes, _ := json.Marshal(a)
	//	fmt.Println(string(bytes))
	//
	//	b := map[string]map[string]interface{}{"1": {"title": "afe1112311Ff", "private": true}, "2": {"title": "afe1112311Ff", "private": true}}
	//	bytes2, _ := json.Marshal(b)
	//	fmt.Println(string(bytes2))
	//
	//
	//	data_map := make(map[string]interface{})
	//	data_map["title"] = "title"
	//
	//	permissions:= make([]map[string]interface{}, 2)
	//	res := make(map[string]interface{})
	//
	//	r:= make([]map[string]interface{}, 1)
	//	inner_r:=make(map[string]interface{})
	//	inner_r["dev_id"] = 424
	//	inner_r["ds_id"] = "fan1"
	//	r[0] = inner_r
	//
	//	res["resources"] = r
	//	permissions[0] =res
	//	permissions[1] =res
	//	data_map["permissions"] = permissions
	//	bytes3, _ := json.Marshal(data_map)
	//	fmt.Println(string(bytes3))

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
func test() {
	smd := smartData.NewSamrtData()
	smd.SetApiKey("2MGfqkx8yTuLA0n9lFBMZLNgGQwA")
	//	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	smd.SetBaseUrl(smartData.DEFAULT_BASE_URL)
	
    datapoints := make(map[string]interface{})
	datapoints["2014-09-01 15:11:01"] = 15
	datapoints["2014-09-01 15:16:01"] = 20
	ret, s := smd.DatapointAdd("66114", "datastream_id1", datapoints)
	if ret == true {
	fmt.Println(ret)
		fmt.Println(*s)
	} else {
		fmt.Println(ret)
		if s != nil {
			fmt.Println(*s)
		}else{
		   fmt.Println(smd.GetErrorNo())
		   fmt.Println(smd.GetError())
		}
	}
	
//	
//    ret, s := smd.DatapointList("66114", "datastream t", nil)
//	if ret == true {
//	fmt.Println(ret)
//		fmt.Println(*s)
//	} else {
//		fmt.Println(ret)
//		if s != nil {
//			fmt.Println(*s)
//		}else{
//		   fmt.Println(smd.GetErrorNo())
//		   fmt.Println(smd.GetError())
//		}
//	}
//	//	smd.SetApiKey(device_key)
//	
//	ret, s = smd.DatapointMultiList("66114", nil)
//	if ret == true {
//	fmt.Println(ret)
//		fmt.Println(*s)
//	} else {
//		fmt.Println(ret)
//		if s != nil {
//			fmt.Println(*s)
//		}else{
//		   fmt.Println(smd.GetErrorNo())
//		   fmt.Println(smd.GetError())
//		}
//	}
}