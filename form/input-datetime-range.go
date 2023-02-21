package form

// 日期时间范围
func InputDatetimeRange(opts ...opt) map[string]interface{} {
	return newForm("input-datetime-range", opts...)
}



// [日期时间选择器值格式](./input-datetime#%E5%80%BC%E6%A0%BC%E5%BC%8F)
func InputDatetimeRange_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// [日期时间选择器显示格式](./input-datetime#%E6%98%BE%E7%A4%BA%E6%A0%BC%E5%BC%8F)
func InputDatetimeRange_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputDatetimeRange_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 日期范围快捷键
func InputDatetimeRange_ranges(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["ranges"] = p
	}
} 

// 限制最小日期时间，用法同 [限制范围](./input-datetime#%E9%99%90%E5%88%B6%E8%8C%83%E5%9B%B4)
func InputDatetimeRange_minDate(p string) opt {
	return func(o map[string]interface{}) {
		o["minDate"] = p
	}
} 

// 限制最大日期时间，用法同 [限制范围](./input-datetime#%E9%99%90%E5%88%B6%E8%8C%83%E5%9B%B4)
func InputDatetimeRange_maxDate(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDate"] = p
	}
} 

// [保存 UTC 值](./input-datetime#utc)
func InputDatetimeRange_utc(p bool) opt {
	return func(o map[string]interface{}) {
		o["utc"] = p
	}
} 

// 是否可清除
func InputDatetimeRange_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 