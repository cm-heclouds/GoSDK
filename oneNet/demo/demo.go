package main

import (
	"fmt"
	"oneNet"
)

func main() {
	on := oneNet.NewOneNet("2MGfqkx8yTuLA0n9lFBMZLNgGQwA")
	datapoints := make(map[string]interface{})
	datapoints["2014-09-01 15:11:01"] = 15
	datapoints["2014-09-01 15:16:01"] = 20
	ret, s := on.DatapointAdd("66114", "datastream_id1", datapoints)
	if ret == true {
		fmt.Println(ret)
		fmt.Println(*s)
	} else {
		fmt.Println(ret)
		if s != nil {
			fmt.Println(*s)
		} else {
			fmt.Println(on.GetErrorNo())
			fmt.Println(on.GetError())
		}
	}

}
