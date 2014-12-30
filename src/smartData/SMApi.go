package smartData

import (
	"bytes"
	"encoding/json"
//	"fmt"
	"io"
	"net/http"
	"strings"
	"net/url"
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
//			fmt.Println(body.(string))
		case map[string]interface{}:
			body_bytes, _ := json.Marshal(body.(map[string]interface{}))
//						fmt.Println(string(body_bytes))
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

	var ret bool = false
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

//	fmt.Println(string(b))
	err := json.Unmarshal(b, &rt_m)

	if err_no, ok := rt_m["errno"]; err == nil && ok {
		if err_no.(float64) == 0 { //no error happened
			if data_map, ok := rt_m["data"]; ok {
				data_byte, _ := json.Marshal(data_map)
				data_str := string(data_byte)
				ret_s = &data_str
			} else {
				data_str := ""
				ret_s = &data_str
			}
			ret = true
		}

		sd.error_no = int(err_no.(float64))
		if err, ok := rt_m["error"]; ok {
			sd.error_ = err.(string)
		}
	} else { //不是json串  或返回有问题
		ret_s = nil
		//暂时。。。。
		sd.error_no = 999
		sd.error_ = "inner error"
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
    val := &url.Values{}
	for k, v := range m {
       val.Set(k, v)
	}
	s =val.Encode()
	if s != "" {
		s = "?" + s
	}
	return
}
