package form

// 日期时间
func InputDatetime(opts ...opt) map[string]interface{} {
	return newForm("input-datetime", opts...)
}



// [默认值](./datetime#%E9%BB%98%E8%AE%A4%E5%80%BC)
func InputDatetime_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
} 

// 日期时间选择器值格式，更多格式类型请参考 [文档](https://momentjs.com/docs/#/displaying/format/)
func InputDatetime_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// 日期时间选择器显示格式，即时间戳格式，更多格式类型请参考 [文档](https://momentjs.com/docs/#/displaying/format/)
func InputDatetime_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 占位文本
func InputDatetime_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 日期时间快捷键
func InputDatetime_shortcuts(p string) opt {
	return func(o map[string]interface{}) {
		o["shortcuts"] = p
	}
} 

// 限制最小日期时间
func InputDatetime_minDate(p string) opt {
	return func(o map[string]interface{}) {
		o["minDate"] = p
	}
} 

// 限制最大日期时间
func InputDatetime_maxDate(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDate"] = p
	}
} 

// 保存 utc 值
func InputDatetime_utc(p bool) opt {
	return func(o map[string]interface{}) {
		o["utc"] = p
	}
} 

// 是否可清除
func InputDatetime_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 是否内联
func InputDatetime_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 

// 请参考 [input-time](./input-time#控制输入范围) 里的说明
func InputDatetime_timeConstraints(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["timeConstraints"] = p
	}
} 