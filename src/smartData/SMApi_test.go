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

func Test_Device(t *testing.T) {
	ret, s := smd.Device(66114)
	t.Log(ret)
	t.Log(*s)
}

func TestX(t *testing.T) {
	//t.Error("vbbb")

}
