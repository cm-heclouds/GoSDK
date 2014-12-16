package smartData

import (
	"encoding/json"
	//	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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
	beforeCall  func(req *http.Request, url, method, body string)
	afterCall   func(req *http.Request, url, method, body string, ret []byte)
	afterDecode func(req *http.Request, url, method, body string, ori_ret []byte, ret bool)
}

func (sd *SmartData) SetApiKey(key string) {
	sd.key = key
}

func (sd *SmartData) SetBaseUrl(base_url string) {
	sd.base_url = base_url
}

func (sd *SmartData) SetAfterCall(fn func(req *http.Request, url, method, body string)) {
	sd.beforeCall = fn
}

func (sd *SmartData) SetBeforeCall(fn func(req *http.Request, url, method, body string, ret []byte)) {
	sd.afterCall = fn
}

func (sd *SmartData) SetAfterDecode(fn func(req *http.Request, url, method, body string, ori_ret []byte, ret bool)) {
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

func (sd *SmartData) Devicelist(dlo *DeviceListOption) (bool, *string) {
	if dlo == nil {
		dlo = DefaultDeviceListOption
	}
	params := make(map[string]string)
	if dlo.page != nil {
		params["page"] = strconv.Itoa(dlo.page.i)
	}

	if dlo.page_size != nil {
		params["page_size"] = strconv.Itoa(dlo.page_size.i)
	}

	if dlo.key_word != nil {
		params["key_word"] = dlo.key_word.s
	}

	if dlo.tag != nil {
		params["tag"] = dlo.tag.s
	}

	if dlo.is_online != nil {
		if dlo.is_online.b == true {
			params["is_online"] = "1"
		} else {
			params["is_online"] = "0"
		}
	}
	api := "/devices?" + pares_params(params)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DeviceAdd(device_json string) (bool, *string) {
	api := "/devices/"
	return sd.call(&api, ALLOW_METHODS["POST"], &device_json, nil)
}

func (sd *SmartData) DeviceEdit(id int, device_json string) (bool, *string) {
	api := "/devices/" + strconv.Itoa(id)
	return sd.call(&api, ALLOW_METHODS["PUT"], &device_json, nil)
}

func (sd *SmartData) DeviceDelete(id int) (bool, *string) {
	api := "/devices/" + strconv.Itoa(id)
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datastream
func (sd *SmartData) Datastream(device_id int, datastream_id int) (bool, *string) {
	api := "/devices/" + strconv.Itoa(device_id) + "/datastreams/" + strconv.Itoa(datastream_id)
	return sd.call(&api, ALLOW_METHODS["GET"], nil, nil)
}

func (sd *SmartData) DatastreamAdd(device_id int, datastream_json string) (bool, *string) {
	api := "/devices/" + strconv.Itoa(device_id) + "/datastreams/"
	return sd.call(&api, ALLOW_METHODS["POST"], &datastream_json, nil)
}

func (sd *SmartData) DatastreamEdit(device_id int, datastream_id int, datastream_json string) (bool, *string) {
	api := "/devices/" + strconv.Itoa(device_id) + "/datastreams/" + strconv.Itoa(datastream_id)
	return sd.call(&api, ALLOW_METHODS["PUT"], &datastream_json, nil)
}

func (sd *SmartData) DatastreamDelete(device_id int, datastream_id int) (bool, *string) {
	api := "/devices/" + strconv.Itoa(device_id) + "/datastreams/" + strconv.Itoa(datastream_id)
	return sd.call(&api, ALLOW_METHODS["DELETE"], nil, nil)
}

//datapoint
func (sd *SmartData) DatapointAdd(device_id int, datastream_id int,) (bool, *string) {
	api := "/devices/"
	return sd.call(&api, ALLOW_METHODS["POST"], nil, nil)
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

func (sd *SmartData) call(url, method, body *string, headers map[string]string) (bool, *string) {
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
		body_reader = strings.NewReader(*body)
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
		sd.beforeCall(req, *url, *method, *body)
	}
	client := &http.Client{}
	resp, _ := client.Do(req)
	b := make([]byte, resp.ContentLength)
	resp.Body.Read(b)
	if sd.afterCall != nil {
		sd.afterCall(req, *url, *method, *body, b)
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
		sd.afterDecode(req, *url, *method, *body, b, ret)
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
