package smartData

import ()

var (
	DefaultDeviceListOption = &DeviceListOption{
		page:      &IntHolder{i: 1},
		page_size: &IntHolder{i: 30},
	}
)

type DeviceListOption struct {
	page, page_size       *IntHolder
	key_word, tag         *StringHolder
	is_online, is_private *BoolHolder
}

func NewDeviceListOption() *DeviceListOption {
	return new(DeviceListOption)
}

func (dlo *DeviceListOption) SetPage(page int) {
	dlo.page = &IntHolder{i: page}
}

func (dlo *DeviceListOption) SetPageSize(page_size int) {
	if page_size > 100 {
		page_size = 100
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
