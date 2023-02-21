package form

// 地理位置
func LocationPicker(opts ...opt) map[string]interface{} {
	return newForm("location-picker", opts...)
}



// 地图厂商，目前只实现了百度地图
func LocationPicker_vendor(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["vendor"] = p
	}
} 

// 百度地图的 ak
func LocationPicker_ak(p string) opt {
	return func(o map[string]interface{}) {
		o["ak"] = p
	}
} 

// 输入框是否可清空
func LocationPicker_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 默认提示
func LocationPicker_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 默为百度坐标，可设置为'gcj02'
func LocationPicker_coordinatesType(p string) opt {
	return func(o map[string]interface{}) {
		o["coordinatesType"] = p
	}
} 