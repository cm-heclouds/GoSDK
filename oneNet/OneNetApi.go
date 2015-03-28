package oneNet

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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

type OneNet struct {
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

func (on *OneNet) SetApiKey(key string) {
	on.key = key
}

func (on *OneNet) SetBaseUrl(base_url string) {
	on.base_url = base_url
}

func (on *OneNet) SetAfterCall(fn func(req *http.Request, url, method string, body interface{})) {
	on.beforeCall = fn
}

func (on *OneNet) SetBeforeCall(fn func(req *http.Request, url, method string, body interface{}, ret []byte)) {
	on.afterCall = fn
}

func (on *OneNet) SetAfterDecode(fn func(req *http.Request, url, method string, body interface{}, ori_ret []byte, ret bool)) {
	on.afterDecode = fn
}

func (on *OneNet) GetApiKey() string {
	return on.key
}

func (on *OneNet) GetHttpCode() int {
	return on.http_code
}

func (on *OneNet) GetErrorNo() int {
	return on.error_no
}

func (on *OneNet) GetError() string {
	return on.error_
}

func (on *OneNet) paddingUrl(url *string) *string {
	if url == nil {
		return nil
	}
	var ret string
	if string((*url)[0]) != "/" {
		ret = on.base_url + "/" + *url
	} else {
		ret = on.base_url + *url
	}
	return &ret
}

func (on *OneNet) call(url, method *string, body interface{}, headers map[string]string) (bool, *string) {
	//check url
	url = on.paddingUrl(url)
	if url == nil {
		on.http_code = 500
		return false, nil
	}

	//check method
	if _, ok := ALLOW_METHODS[*method]; !ok {
		on.http_code = 500
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
	if on.key != "" {
		headers["api-key"] = on.key
	}

	req, _ := http.NewRequest(*method, *url, body_reader)
	for k, v := range headers { //add more header to request
		req.Header.Add(k, v)
	}

	var ret bool = false
	if on.beforeCall != nil {
		on.beforeCall(req, *url, *method, body)
	}
	client := &http.Client{}

	resp, _ := client.Do(req)
	defer resp.Body.Close() //socket文件的关闭...

	b := make([]byte, resp.ContentLength)
	resp.Body.Read(b)
	if on.afterCall != nil {
		on.afterCall(req, *url, *method, body, b)
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

		on.error_no = int(err_no.(float64))
		if err, ok := rt_m["error"]; ok {
			on.error_ = err.(string)
		}
	} else { //不是json串  或返回有问题
		ret_s = nil
		//暂时。。。。
		on.error_no = 999
		on.error_ = "inner error"
	}

	if on.afterDecode != nil {
		on.afterDecode(req, *url, *method, body, b, ret)
	}

	return ret, ret_s
}

func NewOneNet(key string) *OneNet {
	return &OneNet{
		key:       key,
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
	s = val.Encode()
	if s != "" {
		s = "?" + s
	}
	return
}
