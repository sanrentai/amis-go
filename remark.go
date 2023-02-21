package amis

// 标记
func Remark(opts ...opt) map[string]interface{} {
	return newCompent("remark", opts...)
}

// 外层 CSS 类名
func Remark_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 提示文本
func Remark_content(p string) opt {
	return func(o map[string]interface{}) {
		o["content"] = p
	}
}

// 弹出位置
func Remark_placement(p string) opt {
	return func(o map[string]interface{}) {
		o["placement"] = p
	}
}

// 触发条件
func Remark_trigger(p string) opt {
	return func(o map[string]interface{}) {
		o["trigger"] = p
	}
}

// 图标
func Remark_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 图标形状
func Remark_shape(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["shape"] = p
	}
}
