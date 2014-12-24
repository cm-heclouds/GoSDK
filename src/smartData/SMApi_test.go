package smartData

import (
	"testing"
)

var (
	smd *SmartData = &SmartData{
		key:       "2MGfqkx8yTuLA0n9lFBMZLNgGQwA",
		base_url:  DEFAULT_BASE_URL,
		http_code: 200,
		error_no:  0,
		error_:    "",
	}
)

//Teat Pass
//func Test_Device(t *testing.T) {
//	ret, s := smd.Device(66114)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		}
//	}
//}

//Teat Pass
//func TestDeviceList(t *testing.T) {
//	ret, s := smd.DeviceList(nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		}
//	}
//}

//Teat Pass
//func TestDeviceAdd(t *testing.T) {
//	device := make(map[string]interface{})
//	device["title"] = "my test device"
//	device["private"] = true
//	device_key := smd.GetApiKey()
//	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
//	ret, s := smd.DeviceAdd(device)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//	smd.SetApiKey(device_key)
//}

//Teat Pass
//func TestDeviceEdit(t *testing.T) {
//	device := make(map[string]interface{})
//	device["title"] = "test device edited1"
//	device["private"] = true
//	ret, s := smd.DeviceEdit("66114", device)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		}
//	}
//}

//Teat Pass
//func TestDeviceDelete(t *testing.T) {
//	device_key := smd.GetApiKey()
//	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
//	ret, s := smd.DeviceDelete("67310")
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//	smd.SetApiKey(device_key)
//}

//Teat Pass
//func TestDatastream(t *testing.T) {
//	ret, s := smd.Datastream("66114", "temp")
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		}
//	}
//}

//Teat Pass
//func TestDatastreamAdd(t *testing.T) {
//	dataStream := make(map[string]interface{})
//	dataStream["id"] = "datastream t"
//	ret, s := smd.DatastreamAdd("66114", dataStream)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}

//Teat Pass
//func TestDatastreamEdit(t *testing.T) {
//	dataStream := make(map[string]interface{})
//    dataStream["unit"] = "celsius"
//    dataStream["unit_symbol"] = "C"
//	ret, s := smd.DatastreamEdit("66114", "temp", dataStream)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}

//Teat Pass
//func TestDatastreamDelete(t *testing.T) {
//	ret, s := smd.DatastreamDelete("66114", "hum")
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}

//Teat Pass
func TestDatapointAdd(t *testing.T) {
	datapoints := make(map[string]interface{})
	datapoints["2014-09-01 15:06:01"] = 15
	datapoints["2014-09-01 15:09:01"] = 20
	ret, s := smd.DatapointAdd("66114", "datastream_id1", datapoints)
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		} else {
			t.Error(smd.GetErrorNo())
			t.Error(smd.GetError())

		}
	}
}
//
//Teat Pass
//func TestDatapointMultiAdd(t *testing.T) {
//	datapoints := make(map[string]interface{})
//	datapoints["2014-09-01 15:16:01"] = 15
//	datapoints["2014-09-01 15:19:01"] = 20
//
//	datastreams := make(map[string]map[string]interface{})
//	datastreams["temp"] = datapoints
//	datastreams["datastream t"] = datapoints
//
//	ret, s := smd.DatapointMultiAdd("66114", datastreams)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}

//Teat Pass
//func TestDatapointList(t *testing.T) {
//	ret, s := smd.DatapointList("66114", "temp", nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}

//Teat Pass
//func TestDatapointMultiList(t *testing.T) {
//	ret, s := smd.DatapointMultiList("66114", nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}

//func TestDatapointDelete(t *testing.T) {
//	ret, s := smd.DatapointDelete("66114", "datastream_id1", nil, nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}
//
//func TestDatapointMultiDelete(t *testing.T) {
//	ret, s := smd.DatapointMultiDelete("66114", nil, nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//
//		}
//	}
//}
//func TestTriggerAdd(t *testing.T) {
//	trigger := make(map[string]interface{})
//	trigger["url"] = "www.example.com"
//	trigger["type"] = ">"
//	trigger["threshold"] = "threshold "
//
//	ret, s := smd.TriggerAdd("66114", "temp", nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		} else {
//			t.Error(smd.GetErrorNo())
//			t.Error(smd.GetError())
//		}
//	}
//}

//func TestTrigger(t *testing.T) {
//	ret, s := smd.TriggerAdd("66114", "datastream test", nil)
//	if ret == true {
//		t.Log(ret)
//		t.Log(*s)
//	} else {
//		t.Error(ret)
//		if s != nil {
//			t.Error(*s)
//		}
//	}
//}
