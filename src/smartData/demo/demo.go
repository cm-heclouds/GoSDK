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
	smd := smartData.NewSamrtData()
	smd.SetApiKey("2MGfqkx8yTuLA0n9lFBMZLNgGQwA")
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
		} else {
			fmt.Println(smd.GetErrorNo())
			fmt.Println(smd.GetError())
		}
	}

}
