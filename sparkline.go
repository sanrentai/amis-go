package amis

// 走势图
func Sparkline(opts ...opt) map[string]interface{} {
	return newCompent("sparkline", opts...)
}

// 关联的变量
func Sparkline_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 宽度
func Sparkline_width(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}

// 高度
func Sparkline_height(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// 数据为空时显示的内容
func Sparkline_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}
