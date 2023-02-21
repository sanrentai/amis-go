package amis

// Json
func Json(opts ...opt) map[string]interface{} {
	return newCompent("json", opts...)
}

// 外层 CSS 类名
func Json_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// json 值，如果是 string 会自动 parse
func Json_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 通过数据映射获取数据链中的值
func Json_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 占位文本
func Json_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 默认展开的层级
func Json_levelExpand(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["levelExpand"] = p
	}
}

// 主题，可选`twilight`和`eighties`
func Json_jsonTheme(p string) opt {
	return func(o map[string]interface{}) {
		o["jsonTheme"] = p
	}
}

// 是否可修改
func Json_mutable(p bool) opt {
	return func(o map[string]interface{}) {
		o["mutable"] = p
	}
}

// 是否显示数据类型
func Json_displayDataTypes(p bool) opt {
	return func(o map[string]interface{}) {
		o["displayDataTypes"] = p
	}
}

// 设置字符串的最大展示长度，点击字符串可以切换全量/部分展示方式，默认展示全量字符串
func Json_ellipsisThreshold(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["ellipsisThreshold"] = p
	}
}
