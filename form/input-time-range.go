package form

// 时间范围
func InputTimeRange(opts ...opt) map[string]interface{} {
	return newForm("input-time-range", opts...)
}



// [时间范围选择器值格式](./date#%E5%80%BC%E6%A0%BC%E5%BC%8F)
func InputTimeRange_timeFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["timeFormat"] = p
	}
} 

// [时间范围选择器值格式](./date#%E5%80%BC%E6%A0%BC%E5%BC%8F)
func InputTimeRange_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// [时间范围选择器显示格式](./date#%E6%98%BE%E7%A4%BA%E6%A0%BC%E5%BC%8F)
func InputTimeRange_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputTimeRange_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 是否可清除
func InputTimeRange_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 是否内联模式
func InputTimeRange_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 