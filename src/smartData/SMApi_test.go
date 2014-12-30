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

//Test Pass
func Test_Device(t *testing.T) {
	ret, s := smd.Device(66114)
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestDeviceList(t *testing.T) {
	ret, s := smd.DeviceList(nil)
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestDeviceAdd(t *testing.T) {
	device := make(map[string]interface{})
	device["title"] = "my test device"
	device["private"] = true
	device_key := smd.GetApiKey()
	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	ret, s := smd.DeviceAdd(device)
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
	smd.SetApiKey(device_key)
}

//Test Pass
func TestDeviceEdit(t *testing.T) {
	device := make(map[string]interface{})
	device["title"] = "test device edited1"
	device["private"] = true
	ret, s := smd.DeviceEdit("66114", device)
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestDeviceDelete(t *testing.T) {
	device_key := smd.GetApiKey()
	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	ret, s := smd.DeviceDelete("67310")
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
	smd.SetApiKey(device_key)
}

//Test Pass
func TestDatastream(t *testing.T) {
	ret, s := smd.Datastream("66114", "temp")
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestDatastreamAdd(t *testing.T) {
	dataStream := make(map[string]interface{})
	dataStream["id"] = "datastream t"
	ret, s := smd.DatastreamAdd("66114", dataStream)
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

//Test Pass
func TestDatastreamEdit(t *testing.T) {
	dataStream := make(map[string]interface{})
	dataStream["unit"] = "celsius"
	dataStream["unit_symbol"] = "C"
	ret, s := smd.DatastreamEdit("66114", "temp", dataStream)
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

//Test Pass
func TestDatastreamDelete(t *testing.T) {
	ret, s := smd.DatastreamDelete("66114", "hum")
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

//Test Pass
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

//Test Pass
func TestDatapointMultiAdd(t *testing.T) {
	datapoints := make(map[string]interface{})
	datapoints["2014-09-01 15:16:01"] = 15
	datapoints["2014-09-01 15:19:01"] = 20

	datastreams := make(map[string]map[string]interface{})
	datastreams["temp"] = datapoints
	datastreams["datastream t"] = datapoints

	ret, s := smd.DatapointMultiAdd("66114", datastreams)
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

//Test Pass
func TestDatapointList(t *testing.T) {
	ret, s := smd.DatapointList("66114", "temp", nil)
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

//Test Pass
func TestDatapointMultiList(t *testing.T) {
	ret, s := smd.DatapointMultiList("66114", nil)
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

//Test Pass
func TestDatapointDelete(t *testing.T) {
	ret, s := smd.DatapointDelete("66114", "datastream t", "2011-01-02 15:04:02", nil)
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

//Test Pass
func TestDatapointMultiDelete(t *testing.T) {
	ret, s := smd.DatapointMultiDelete("66114", "2011-01-02 15:04:02", nil)
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

//Test Pass
func TestTriggerAdd(t *testing.T) {
	trigger := make(map[string]interface{})
	trigger["url"] = "www.example.com"
	trigger["type"] = ">"
	trigger["threshold"] = 100

	ret, s := smd.TriggerAdd("66114", "temp", trigger)
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

//Test Pass
func TestTrigger(t *testing.T) {
	ret, s := smd.Trigger("66114", "temp", "10811")
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestTriggerEdit(t *testing.T) {
	trigger := make(map[string]interface{})
	trigger["url"] = "www.example.comaaaaaaaaaaaaaaaa"
	trigger["type"] = ">"
	trigger["threshold"] = 100

	ret, s := smd.TriggerEdit("66114", "temp", "10810", trigger)
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestTriggerDelete(t *testing.T) {
	ret, s := smd.TriggerDelete("66114", "temp", "10811")
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
}

//Test Pass
func TestApiKey(t *testing.T) {
	device_key := smd.GetApiKey()
	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	ret, s := smd.ApiKey("66114")
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
	smd.SetApiKey(device_key)
}

//Test Pass
func TestApiKeyAdd(t *testing.T) {
	device_key := smd.GetApiKey()
	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	dev_ids := []string{"66114"}
	ret, s := smd.ApiKeyAdd(dev_ids, "api_key test")
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
	smd.SetApiKey(device_key)
}

//Test Pass
func TestApiKeyDelete(t *testing.T) {
	device_key := smd.GetApiKey()
	smd.SetApiKey("gJNoxz2hn1nPa3WdZkmVdUu2Ow4A")
	ret, s := smd.ApiKeyDelete("NV7Xtt8onqXVxorbS1q2FCxW9KcA")
	if ret == true {
		t.Log(ret)
		t.Log(*s)
	} else {
		t.Error(ret)
		if s != nil {
			t.Error(*s)
		}
	}
	smd.SetApiKey(device_key)
}
