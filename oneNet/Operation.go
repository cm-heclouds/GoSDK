package oneNet

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

//设备相关API
func (on *OneNet) Device(id int) (bool, *string) {
	api := "/devices/" + strconv.Itoa(id)
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (on *OneNet) DeviceList(dlo *DeviceListOption) (bool, *string) {
	if dlo == nil {
		dlo = DefaultDeviceListOption
	}
	params := make(map[string]string)
	parseOption(dlo, params)
	api := "/devices" + pares_params(params)
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (on *OneNet) DeviceAdd(device interface{}) (bool, *string) {
	api := "/devices"
	return on.call(&api, ALLOW_METHODS["POST"], device, nil)
}

func (on *OneNet) DeviceEdit(id string, device interface{}) (bool, *string) {
	api := "/devices/" + id
	return on.call(&api, ALLOW_METHODS["PUT"], device, nil)
}

func (on *OneNet) DeviceDelete(id string) (bool, *string) {
	api := "/devices/" + id
	return on.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datastream
func (on *OneNet) Datastream(device_id, datastream_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (on *OneNet) DatastreamAdd(device_id string, datastream interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams"
	return on.call(&api, ALLOW_METHODS["POST"], datastream, nil)
}

func (on *OneNet) DatastreamEdit(device_id, datastream_id string, datastream interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return on.call(&api, ALLOW_METHODS["PUT"], datastream, nil)
}

func (on *OneNet) DatastreamDelete(device_id, datastream_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return on.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datapoint
/*
  datapoint:   array (timestamp -> value)
    1. map[timestamp] value
    2. []string{"timestamp:value",}
  timestamp :   year-month-day hour:minute:second
*/
func (on *OneNet) DatapointAdd(device_id, datastream_id string, datapoint interface{}) (bool, *string) {
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
	//	data_bytes, _ := json.Marshal(data_m)

	return on.call(&api, ALLOW_METHODS["POST"], data_m, nil)
}

/*
  data:   array (datastream_id->array (timestamp[year:month:day hour:minute:second] -> value))
      map[string]map[timestamp]value
*/
func (on *OneNet) DatapointMultiAdd(device_id string, datas map[string]map[string]interface{}) (bool, *string) {
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
	//	data_bytes, _ := json.Marshal(data_m)

	return on.call(&api, ALLOW_METHODS["POST"], data_m, nil)
}

func (on *OneNet) DatapointList(device_id, datastream_id string, dplo *DataPointListOption) (bool, *string) {
	if dplo == nil {
		dplo = DefaultDataPointListOption
	}

	params := make(map[string]string)
	params["datastream_id"] = datastream_id
	parseOption(dplo, params)
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (on *OneNet) DatapointMultiList(device_id string, dplo *DataPointListOption) (bool, *string) {
	if dplo == nil {
		dplo = DefaultDataPointListOption
	}
	params := make(map[string]string)
	parseOption(dplo, params)
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

/*
   start_time,end_time:
      1. string type : year-month-day hour:minute:second
      2. time.Time or *time.Time type
*/
func (on *OneNet) DatapointDelete(device_id, datastream_id string, start_time, end_time interface{}) (bool, *string) {
	params := make(map[string]string)

	if start_time != nil {
		stime := new(time.Time)
		parseTime(start_time, stime)
		params["start"] = stime.Format("2006-01-02T15:04:02")
	}

	if end_time != nil {
		etime := new(time.Time)
		parseTime(end_time, etime)
		params["end"] = etime.Format("2006-01-02T15:04:02")
	}

	//	if start_time != nil && end_time != nil {
	//		etime := new(time.Time)
	//
	//
	//
	//		params["duration"] = strconv.Itoa(int(etime.Sub(*stime).Seconds()))
	//	}
	params["datastream_id"] = datastream_id
	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return on.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

func (on *OneNet) DatapointMultiDelete(device_id string, start_time, end_time interface{}) (bool, *string) {
	params := make(map[string]string)
	//	if start_time != nil && end_time != nil {
	//		etime := new(time.Time)
	//		stime := new(time.Time)
	//		parseTime(start_time, stime)
	//		parseTime(end_time, etime)
	//		params["start"] = stime.Format("2006-01-02T15:04:02")
	//		params["duration"] = strconv.Itoa(int(etime.Sub(*stime).Seconds()))
	//	}
	if start_time != nil {
		stime := new(time.Time)
		parseTime(start_time, stime)
		params["start"] = stime.Format("2006-01-02T15:04:02")
	}

	if end_time != nil {
		etime := new(time.Time)
		parseTime(end_time, etime)
		params["end"] = etime.Format("2006-01-02T15:04:02")
	}

	api := "/devices/" + device_id + "/datapoints" + pares_params(params)
	return on.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

func (on *OneNet) Trigger(device_id, datastream_id, trigger_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers/" + trigger_id
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (on *OneNet) TriggerAdd(device_id, datastream_id string, trigger interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers"
	return on.call(&api, ALLOW_METHODS["POST"], trigger, nil)
}

func (on *OneNet) TriggerEdit(device_id, datastream_id, trigger_id string, trigger interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers/" + trigger_id
	return on.call(&api, ALLOW_METHODS["PUT"], trigger, nil)
}

func (on *OneNet) TriggerDelete(device_id, datastream_id, trigger_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id + "/triggers/" + trigger_id
	return on.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//获取APIkey
func (on *OneNet) ApiKey(device_id string) (bool, *string) {
	v := &url.Values{}
	v.Set("dev_id", device_id)
	api := "/keys?" + v.Encode()
	return on.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

//暂时只提供到dev_id级别权限, 必须是master_key 才行
func (on *OneNet) ApiKeyAdd(device_ids []string, title string) (bool, *string) {
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

	//	data_bytes, _ := json.Marshal(data_map)

	return on.call(&api, ALLOW_METHODS["POST"], data_map, nil)
}

//key_string 为需要删除的key的string
func (on *OneNet) ApiKeyDelete(key_string string) (bool, *string) {
	api := "/keys/" + url.QueryEscape(key_string)
	return on.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}
