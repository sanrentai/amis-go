package form

// 月份
func InputMonth(opts ...opt) map[string]interface{} {
	return newForm("input-month", opts...)
}



// [默认值](./date#%E9%BB%98%E8%AE%A4%E5%80%BC)
func InputMonth_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
} 

// 月份选择器值格式，更多格式类型请参考 [moment](http://momentjs.com/)
func InputMonth_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// 月份选择器显示格式，即时间戳格式，更多格式类型请参考 [moment](http://momentjs.com/)
func InputMonth_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputMonth_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 是否可清除
func InputMonth_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 