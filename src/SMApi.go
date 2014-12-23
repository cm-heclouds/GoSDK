package smartData

import (
	"encoding/json"
	//	"fmt"
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_BASE_URL = "http://api.heclouds.com/"
)

var (
	get_          = "GET"
	put_          = "PUT"
	post_         = "POST"
	delete_       = "DELETE"
	ALLOW_METHODS = map[string]*string{"GET": &get_, "PUT": &put_, "POST": &post_, "DELETE": &delete_}
)

type SmartData struct {
	key       string
	base_url  string
	http_code int
	error_no  int
	error_    string

	//自定义调用的函数
	beforeCall  func(req *http.Request, url, method string, body interface{})
	afterCall   func(req *http.Request, url, method string, body interface{}, ret []byte)
	afterDecode func(req *http.Request, url, method string, body interface{}, ori_ret []byte, ret bool)
}

func (sd *SmartData) SetApiKey(key string) {
	sd.key = key
}

func (sd *SmartData) SetBaseUrl(base_url string) {
	sd.base_url = base_url
}

func (sd *SmartData) SetAfterCall(fn func(req *http.Request, url, method string, body interface{})) {
	sd.beforeCall = fn
}

func (sd *SmartData) SetBeforeCall(fn func(req *http.Request, url, method string, body interface{}, ret []byte)) {
	sd.afterCall = fn
}

func (sd *SmartData) SetAfterDecode(fn func(req *http.Request, url, method string, body interface{}, ori_ret []byte, ret bool)) {
	sd.afterDecode = fn
}

func (sd *SmartData) GetApiKey() string {
	return sd.key
}

func (sd *SmartData) GetHttpCode() int {
	return sd.http_code
}

func (sd *SmartData) GetErrorNo() int {
	return sd.error_no
}

func (sd *SmartData) GetError() string {
	return sd.error_
}

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
	api := "/devices?" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DeviceAdd(device interface{}) (bool, *string) {
	api := "/devices/"
	return sd.call(&api, ALLOW_METHODS["POST"], device, nil)
}

func (sd *SmartData) DeviceEdit(id int, device interface{}) (bool, *string) {
	api := "/devices/" + strconv.Itoa(id)
	return sd.call(&api, ALLOW_METHODS["PUT"], device, nil)
}

func (sd *SmartData) DeviceDelete(id int) (bool, *string) {
	api := "/devices/" + strconv.Itoa(id)
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datastream
func (sd *SmartData) Datastream(device_id, datastream_id string) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + datastream_id
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DatastreamAdd(device_id string, datastream interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/"
	return sd.call(&api, ALLOW_METHODS["POST"], datastream, nil)
}

func (sd *SmartData) DatastreamEdit(device_id string, datastream_id int, datastream interface{}) (bool, *string) {
	api := "/devices/" + device_id + "/datastreams/" + strconv.Itoa(datastream_id)
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

	data_bytes, _ := json.Marshal(data_map)

	datapoint_str := `{"datastreams":[` + string(data_bytes) + `,]}`

	return sd.call(&api, ALLOW_METHODS["POST"], datapoint_str, nil)
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
	api := "/devices/" + device_id + "/datapoints?" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DatapointMultiList(device_id string, dplo *DataPointListOption) (bool, *string) {
	if dplo == nil {
		dplo = DefaultDataPointListOption
	}
	params := make(map[string]string)
	parseOption(dplo, params)
	api := "/devices/" + device_id + "/datapoints?" + pares_params(params)
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
	api := "/devices/" + device_id + "/datapoints?" + pares_params(params)
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
	api := "/devices/" + device_id + "/datapoints?" + pares_params(params)
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

func (sd *SmartData) ApiKey(device_id string) (bool, *string) {
	api := "/keys?dev_id=" + device_id
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) paddingUrl(url *string) *string {
	if url == nil {
		return nil
	}
	var ret string
	if string((*url)[0]) != "/" {
		ret = sd.base_url + "/" + *url
	} else {
		ret = sd.base_url + *url
	}
	return &ret
}

func (sd *SmartData) call(url, method *string, body interface{}, headers map[string]string) (bool, *string) {
	//check url
	url = sd.paddingUrl(url)
	if url == nil {
		sd.http_code = 500
		return false, nil
	}

	//check method
	if _, ok := ALLOW_METHODS[*method]; !ok {
		sd.http_code = 500
		return false, nil
	}

	//check body
	var body_reader io.Reader
	if body != nil {
		switch body.(type) {
		case string:
			body_reader = strings.NewReader(body.(string))
		case map[string]interface{}:
			body_bytes, _ := json.Marshal(body.(map[string]interface{}))
			body_reader = bytes.NewReader(body_bytes)
		case []string:
			m := make(map[string]interface{})
			for _, s := range body.([]string) {
				part := strings.SplitN(":", s, 2)
				if len(part) == 2 {
					m[part[0]] = part[1]
				}
			}
			body_bytes, _ := json.Marshal(m)
			body_reader = bytes.NewReader(body_bytes)
		default:
			body_reader = nil
		}

	} else {
		body_reader = nil
	}

	//check header
	if headers == nil {
		headers = make(map[string]string)
	}

	//add api-key to headers
	if sd.key != "" {
		headers["api-key"] = sd.key
	}

	req, _ := http.NewRequest(*method, *url, body_reader)
	for k, v := range headers { //add more header to request
		req.Header.Add(k, v)
	}

	var ret bool = true
	if sd.beforeCall != nil {
		sd.beforeCall(req, *url, *method, body)
	}
	client := &http.Client{}
	resp, _ := client.Do(req)
	b := make([]byte, resp.ContentLength)
	resp.Body.Read(b)
	if sd.afterCall != nil {
		sd.afterCall(req, *url, *method, body, b)
	}

	var ret_s *string
	var rt_m map[string]interface{}
	err := json.Unmarshal(b, &rt_m)

	if err != nil { //不是json串
		ret = false
		ret_s = nil
	} else {
		if err_no, ok := rt_m["errno"]; ok {
			if err_no.(float64) == 0 { //no error happened
				if data_map, ok := rt_m["data"]; ok {
					data_byte, _ := json.Marshal(data_map)
					data_str := string(data_byte)
					ret_s = &data_str
				}
			} else {
				ret = false
			}
			sd.error_no = int(err_no.(float64))
			if err, ok := rt_m["error"]; ok {
				sd.error_ = err.(string)
			}
		}
	}

	if sd.afterDecode != nil {
		sd.afterDecode(req, *url, *method, body, b, ret)
	}

	return ret, ret_s
}

func NewSamrtData() *SmartData {
	return &SmartData{
		key:       "",
		base_url:  DEFAULT_BASE_URL,
		http_code: 200,
		error_no:  0,
		error_:    "",
	}
}

func pares_params(m map[string]string) (s string) {
	length := len(m)
	count := 0
	s = ""
	for k, v := range m {
		count++
		if count == length {
			s = s + k + "=" + v
		} else {
			s = s + k + "=" + v + "&"
		}
	}
	return
}
