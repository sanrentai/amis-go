package amis

// 日期时间
func Date(opts ...opt) map[string]interface{} {
	return newCompent("date", opts...)
}

// 外层 CSS 类名
func Date_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 显示的日期数值
func Date_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 在其他组件中，时，用作变量映射
func Date_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 占位内容
func Date_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 展示格式, 更多格式类型请参考 [文档](https://momentjs.com/docs/#/displaying/format/)
func Date_format(p string) opt {
	return func(o map[string]interface{}) {
		o["format"] = p
	}
}

// 数据格式，默认为时间戳。更多格式类型请参考 [文档](https://momentjs.com/docs/#/displaying/format/)
func Date_valueFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["valueFormat"] = p
	}
}

// 是否显示相对当前的时间描述，比如: 11 小时前、3 天前、1 年前等，fromNow 为 true 时，format 不生效。
func Date_fromNow(p bool) opt {
	return func(o map[string]interface{}) {
		o["fromNow"] = p
	}
}

// 更新频率， 默认为 1 分钟
func Date_updateFrequency(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["updateFrequency"] = p
	}
}
