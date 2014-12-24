package smartData

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//设备相关API
func (sd *SmartData) Device(id int) (bool, *string) {
	api := "/devices/" + strconv.Itoa(id)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DeviceList(dlo *DeviceListOption) (bool, *string) {
	if dlo == nil {
		dlo = DefaultDeviceListOption
	}
	params := make(map[string]string)
	parseOption(dlo, params)
	api := "/devices" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DeviceAdd(device interface{}) (bool, *string) {
	api := "/devices"
	return sd.call(&api, ALLOW_METHODS["POST"], device, nil)
}

func (sd *SmartData) DeviceEdit(id string, device interface{}) (bool, *string) {
	api := "/devices/" + id
	return sd.call(&api, ALLOW_METHODS["PUT"], device, nil)
}

func (sd *SmartData) DeviceDelete(id string) (bool, *string) {
	api := "/devices/" + id
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datastream
func (sd *SmartData) Datastream(device_id, datastream_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DatastreamAdd(device_id string, datastream interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams"
	return sd.call(&api, ALLOW_METHODS["POST"], datastream, nil)
}

func (sd *SmartData) DatastreamEdit(device_id, datastream_id string, datastream interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return sd.call(&api, ALLOW_METHODS["PUT"], datastream, nil)
}

func (sd *SmartData) DatastreamDelete(device_id, datastream_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datapoint
/*
  datapoint:   array (timestamp -> value)
    1. map[timestamp] value
    2. []string{"timestamp:value",}
  timestamp :   year-month-day hour:minute:second
*/
func (sd *SmartData) DatapointAdd(device_id, datastream_id string, datapoint interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datapoints"
	var datapoint_maps []map[string]interface{}
	switch datapoint.(type) {
	case []string:
		datapoint_maps = make([]map[string]interface{}, len(datapoint.([]string)))
		for i, s := range datapoint.([]string) {
			m := make(map[string]interface{})
			part := strings.SplitN(":", s, 2)
			if len(part) == 2 {
				tfd, _ := time.Parse("2006-01-02 15:04:02", part[0])
				m["at"] = tfd.Format("2006-01-02T15:04:02")
				m["value"] = part[1]
			}
			datapoint_maps[i] = m
		}
	case map[string]interface{}:
		datapoint_maps = make([]map[string]interface{}, len(datapoint.(map[string]interface{})))
		count := 0
		for k, v := range datapoint.(map[string]interface{}) {
			m := make(map[string]interface{})
			tfd, _ := time.Parse("2006-01-02 15:04:02", k)
			m["at"] = tfd.Format("2006-01-02T15:04:02")
			m["value"] = v
			datapoint_maps[count] = m
			count++
		}
	default:
		datapoint_maps = nil
	}

	data_map := make(map[string]interface{})
	data_map["id"] = datastream_id
	data_map["datapoints"] = datapoint_maps

	multi_data := make([]interface{}, 1)
	multi_data[0] = data_map

	data_m := make(map[string]interface{})
	data_m["datastreams"] = multi_data
	data_bytes, _ := json.Marshal(data_m)

	return sd.call(&api, ALLOW_METHODS["POST"], string(data_bytes), nil)
}

/*
  data:   array (datastream_id->array (timestamp[year:month:day hour:minute:second] -> value))
      map[string]map[timestamp]value
*/
func (sd *SmartData) DatapointMultiAdd(device_id string, datas map[string]map[string]interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datapoints"
	var multi_data []interface{} = make([]interface{}, len(datas))
	pos := 0
	for id, data := range datas {
		datapoint_maps := make([]map[string]interface{}, len(data))
		count := 0
		for k, v := range data {
			m := make(map[string]interface{})
			tfd, _ := time.Parse("2006-01-02 15:04:02", k)
			m["at"] = tfd.Format("2006-01-02T15:04:02")
			m["value"] = v
			datapoint_maps[count] = m
			count++
		}
		data_map := make(map[string]interface{})
		data_map["id"] = id
		data_map["datapoints"] = datapoint_maps
		multi_data[pos] = data_map
		pos++
	}

	data_m := make(map[string]interface{})
	data_m["datastreams"] = multi_data
	data_bytes, _ := json.Marshal(data_m)

	return sd.call(&api, ALLOW_METHODS["POST"], string(data_bytes), nil)
}

func (sd *SmartData) DatapointList(device_id, datastream_id string, dplo *DataPointListOption) (bool, *string) {
	if dplo == nil {
		dplo = DefaultDataPointListOption
	}
	
	params := make(map[string]string)
	params["datastream_id"] = datastream_id
	parseOption(dplo, params)
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DatapointMultiList(device_id string, dplo *DataPointListOption) (bool, *string) {
	if dplo == nil {
		dplo = DefaultDataPointListOption
	}
	params := make(map[string]string)
	parseOption(dplo, params)
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DatapointDelete(device_id, datastream_id string, start_time, end_time interface{}) (bool, *string) {
	params := make(map[string]string)
	if start_time != nil && end_time != nil {
		etime := new(time.Time)
		stime := new(time.Time)
		parseTime(start_time, stime)
		parseTime(end_time, etime)
		params["start"] = stime.Format("2006-01-02T15:04:02")
		params["duration"] = strconv.Itoa(int(etime.Sub(*stime).Seconds()))
	}
	params["datastream_id"] = datastream_id
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

func (sd *SmartData) DatapointMultiDelete(device_id string, start_time, end_time interface{}) (bool, *string) {
	params := make(map[string]string)
	if start_time != nil && end_time != nil {
		etime := new(time.Time)
		stime := new(time.Time)
		parseTime(start_time, stime)
		parseTime(end_time, etime)
		params["start"] = stime.Format("2006-01-02T15:04:02")
		params["duration"] = strconv.Itoa(int(etime.Sub(*stime).Seconds()))
	}
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

func (sd *SmartData) Trigger(device_id, datastream_id, trigger_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers/" + trigger_id
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) TriggerAdd(device_id, datastream_id string, trigger interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers"
	return sd.call(&api, ALLOW_METHODS["POST"], trigger, nil)
}

func (sd *SmartData) TriggerEdit(device_id, datastream_id, trigger_id string, trigger interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers/" + trigger_id
	return sd.call(&api, ALLOW_METHODS["PUT"], trigger, nil)
}

func (sd *SmartData) TriggerDelete(device_id, datastream_id, trigger_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers/" + trigger_id
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//获取APIkey
func (sd *SmartData) ApiKey(device_id string) (bool, *string) {
	v := &url.Values{}
	v.Set("dev_id", device_id)
	api := "/keys?" + v.Encode()
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

//暂时只提供到dev_id级别权限, 必须是master_key 才行
func (sd *SmartData) ApiKeyAdd(device_ids []string, title string) (bool, *string) {
	api := "/keys"
	permissions := make([]map[string]interface{}, len(device_ids))
	for id, device_id := range device_ids {
		res := make(map[string]interface{})
		r := make([]map[string]interface{}, 1)
		inner_r := make(map[string]interface{})
		inner_r["dev_id"] = device_id
		r[0] = inner_r
		res["resources"] = r
		permissions[id] = res
	}

	data_map := make(map[string]interface{})
	data_map["title"] = title
	data_map["permissions"] = permissions

	data_bytes, _ := json.Marshal(data_map)

	return sd.call(&api, ALLOW_METHODS["POST"], string(data_bytes), nil)
}

func (sd *SmartData) ApiKeyDelete(device_id string) (bool, *string) {
	v := &url.Values{}
	v.Set("dev_id", device_id)
	api := "/keys?" + v.Encode()
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}
