package form

// 日期范围
func InputDateRange(opts ...opt) map[string]interface{} {
	return newForm("input-date-range", opts...)
}



// [日期选择器值格式](./date#%E5%80%BC%E6%A0%BC%E5%BC%8F)
func InputDateRange_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// [日期选择器显示格式](./date#%E6%98%BE%E7%A4%BA%E6%A0%BC%E5%BC%8F)
func InputDateRange_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputDateRange_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 日期范围快捷键
func InputDateRange_ranges(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["ranges"] = p
	}
} 

// 限制最小日期，用法同 [限制范围](./date#%E9%99%90%E5%88%B6%E8%8C%83%E5%9B%B4)
func InputDateRange_minDate(p string) opt {
	return func(o map[string]interface{}) {
		o["minDate"] = p
	}
} 

// 限制最大日期，用法同 [限制范围](./date#%E9%99%90%E5%88%B6%E8%8C%83%E5%9B%B4)
func InputDateRange_maxDate(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDate"] = p
	}
} 

// 限制最小跨度，如： 2days
func InputDateRange_minDuration(p string) opt {
	return func(o map[string]interface{}) {
		o["minDuration"] = p
	}
} 

// 限制最大跨度，如：1year
func InputDateRange_maxDuration(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDuration"] = p
	}
} 

// [保存 UTC 值](./date#utc)
func InputDateRange_utc(p bool) opt {
	return func(o map[string]interface{}) {
		o["utc"] = p
	}
} 

// 是否可清除
func InputDateRange_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 是否内联模式
func InputDateRange_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 