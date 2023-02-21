package form

// 日期
func InputDate(opts ...opt) map[string]interface{} {
	return newForm("input-date", opts...)
}



// [默认值](./date#%E9%BB%98%E8%AE%A4%E5%80%BC)
func InputDate_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
} 

// 日期选择器值格式，更多格式类型请参考 [文档](https://momentjs.com/docs/#/displaying/format/)
func InputDate_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
} 

// 日期选择器显示格式，即时间戳格式，更多格式类型请参考 [文档](https://momentjs.com/docs/#/displaying/format/)
func InputDate_inputFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["inputFormat"] = p
	}
} 

// 点选日期后，是否马上关闭选择框
func InputDate_closeOnSelect(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnSelect"] = p
	}
} 

// 占位文本
func InputDate_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
} 

// 日期快捷键
func InputDate_shortcuts(p string) opt {
	return func(o map[string]interface{}) {
		o["shortcuts"] = p
	}
} 

// 限制最小日期
func InputDate_minDate(p string) opt {
	return func(o map[string]interface{}) {
		o["minDate"] = p
	}
} 

// 限制最大日期
func InputDate_maxDate(p string) opt {
	return func(o map[string]interface{}) {
		o["maxDate"] = p
	}
} 

// 保存 utc 值
func InputDate_utc(p bool) opt {
	return func(o map[string]interface{}) {
		o["utc"] = p
	}
} 

// 是否可清除
func InputDate_clearable(p bool) opt {
	return func(o map[string]interface{}) {
		o["clearable"] = p
	}
} 

// 是否内联模式
func InputDate_embed(p bool) opt {
	return func(o map[string]interface{}) {
		o["embed"] = p
	}
} 