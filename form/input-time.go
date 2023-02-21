package form

// 时间
func InputTime(opts ...opt) map[string]interface{} {
	return newForm("input-time", opts...)
}



// [默认值](./date#%E9%BB%98%E8%AE%A4%E5%80%BC)
func InputTime_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
} 

// 时间选择器值格式，更多格式类型请参考 [moment](http://momentjs.com/)
func InputTime_timeFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["timeFormat"] = p
	}
} 

// 时间选择器值格式，更多格式类型请参考 [moment](http://momentjs.com/)
func InputTime_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// 时间选择器显示格式，即时间戳格式，更多格式类型请参考 [moment](http://momentjs.com/)
func InputTime_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputTime_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 是否可清除
func InputTime_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 
func InputTime_timeConstraints(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["timeConstraints"] = p
	}
} 