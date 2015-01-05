GoSDK
=====

设备云平台主要提供Restful方式的接口供开发者调用。
接入设备云请进入 [设备云主站](http://www.heclouds.com) 了解相关文档。

**传送门**:
[API开发文档](http://www.heclouds.com/develop/doc/api/restfullist)

API调用基础地址为:
`http://api.heclouds.com`


简单示例：
```
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
```
