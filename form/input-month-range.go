package form

// 月份范围
func InputMonthRange(opts ...opt) map[string]interface{} {
	return newForm("input-month-range", opts...)
}



// [日期选择器值格式](./date#%E5%80%BC%E6%A0%BC%E5%BC%8F)
func InputMonthRange_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// [日期选择器显示格式](./date#%E6%98%BE%E7%A4%BA%E6%A0%BC%E5%BC%8F)
func InputMonthRange_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputMonthRange_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 限制最小日期，用法同 [限制范围](./date#%E9%99%90%E5%88%B6%E8%8C%83%E5%9B%B4)
func InputMonthRange_minDate(p string) opt {
	return func(o map[string]interface{}) {
		o["minDate"] = p
	}
} 

// 限制最大日期，用法同 [限制范围](./date#%E9%99%90%E5%88%B6%E8%8C%83%E5%9B%B4)
func InputMonthRange_maxDate(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDate"] = p
	}
} 

// 限制最小跨度，如： 2days
func InputMonthRange_minDuration(p string) opt {
	return func(o map[string]interface{}) {
		o["minDuration"] = p
	}
} 

// 限制最大跨度，如：1year
func InputMonthRange_maxDuration(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDuration"] = p
	}
} 

// [保存 UTC 值](./date#utc)
func InputMonthRange_utc(p bool) opt {
	return func(o map[string]interface{}) {
		o["utc"] = p
	}
} 

// 是否可清除
func InputMonthRange_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 是否内联模式
func InputMonthRange_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 