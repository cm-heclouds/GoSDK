package oneNet

import (
	"strconv"
	"time"
)

var (
	DefaultDeviceListOption = &DeviceListOption{
		page:      &IntHolder{i: 1},
		page_size: &IntHolder{i: 30},
	}

	DefaultDataPointListOption = &DataPointListOption{
		page:       &IntHolder{i: 1},
		page_size:  &IntHolder{i: 30},
		order_desc: &BoolHolder{b: false},
	}
)

type StringHolder struct {
	s string
}

type IntHolder struct {
	i int
}

type BoolHolder struct {
	b bool
}

type DeviceListOption struct {
	page, page_size       *IntHolder
	key_word, tag         *StringHolder
	is_online, is_private *BoolHolder
}

func NewDeviceListOption() *DeviceListOption {
	return new(DeviceListOption)
}

func (dlo *DeviceListOption) SetPage(page int) {
	if page < 1 {
		page = 1
	}
	dlo.page = &IntHolder{i: page}
}

func (dlo *DeviceListOption) SetPageSize(page_size int) {
	if page_size > 100 {
		page_size = 100
	}
	if page_size < 1 {
		page_size = 1
	}
	dlo.page_size = &IntHolder{i: page_size}
}

func (dlo *DeviceListOption) SetKeyWord(key_word string) {
	dlo.key_word = &StringHolder{s: key_word}
}

func (dlo *DeviceListOption) SetTag(tag string) {
	dlo.tag = &StringHolder{s: tag}
}

func (dlo *DeviceListOption) SetOnline(is_online bool) {
	dlo.is_online = &BoolHolder{b: is_online}
}

func (dlo *DeviceListOption) SetPrivate(is_private bool) {
	dlo.is_private = &BoolHolder{b: is_private}
}

type DataPointListOption struct {
	start_time, end_time *time.Time
	page, page_size      *IntHolder
	order_desc           *BoolHolder
}

func (dplo *DataPointListOption) SetStartTime(start_time interface{}) {
	parseTime(start_time, dplo.start_time)
}

func (dplo *DataPointListOption) SetEndTime(end_time interface{}) {
	parseTime(end_time, dplo.end_time)
}

func (dplo *DataPointListOption) SetPage(page int) {
	if page < 1 {
		page = 1
	}
	dplo.page = &IntHolder{i: page}
}

func (dplo *DataPointListOption) SetPageSize(page_size int) {
	if page_size > 1000 {
		page_size = 1000
	}
	if page_size < 1 {
		page_size = 1
	}
	dplo.page_size = &IntHolder{i: page_size}
}

func (dplo *DataPointListOption) SetOrderDesc(order_desc bool) {
	dplo.order_desc = &BoolHolder{b: order_desc}
}

func NewDataPointListOption() *DataPointListOption {
	return new(DataPointListOption)
}

//解析目标时间为Time对象，src_time为源，dest_time指向解析后目标对象
func parseTime(src_time interface{}, dest_time *time.Time) {
	switch src_time.(type) {
	case string:
		tfd, _ := time.Parse("2006-01-02 15:04:02", src_time.(string))
		dest_time = &tfd
	case time.Time:
		t := src_time.(time.Time)
		dest_time = &t
	case *time.Time:
		dest_time = src_time.(*time.Time)
	default:
		panic("time parese error,wrong type")
	}
}

func parseOption(option interface{}, params map[string]string) {
	switch option.(type) {
	case *DeviceListOption:
		dlo := option.(*DeviceListOption)
		if dlo.page != nil {
			params["page"] = strconv.Itoa(dlo.page.i)
		}

		if dlo.page_size != nil {
			params["per_page"] = strconv.Itoa(dlo.page_size.i)
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
	case *DataPointListOption:
		dplo := option.(*DataPointListOption)
		if dplo.page != nil {
			params["page"] = strconv.Itoa(dplo.page.i)
		}

		if dplo.page_size != nil {
			params["per_page"] = strconv.Itoa(dplo.page_size.i)
		}

		if dplo.order_desc != nil {
			if dplo.order_desc.b {
				params["sort_time"] = "-1"
			}
		}

		if dplo.start_time != nil {
			params["start"] = dplo.start_time.Format("2006-01-02T15:04:02")
			if dplo.end_time != nil {
				params["duration"] = strconv.Itoa(int(dplo.end_time.Sub(*dplo.start_time).Seconds()))
			} else {
				params["duration"] = strconv.Itoa(int(time.Now().Sub(*dplo.start_time).Seconds()))
			}
		}
	default:
		panic("wrong option type")
	}
}
