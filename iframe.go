package amis

// iFrame
func iFrame(opts ...opt) map[string]interface{} {
	return newCompent("iframe", opts...)
}

// iFrame 的类名
func iFrame_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// frameBorder
func iFrame_frameBorder(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["frameBorder"] = p
	}
}

// 样式对象
func iFrame_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}

// iframe 地址
func iFrame_src(p string) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}

// allow 配置
func iFrame_allow(p string) opt {
	return func(o map[string]interface{}) {
		o["allow"] = p
	}
}

// sandbox 配置
func iFrame_sandbox(p string) opt {
	return func(o map[string]interface{}) {
		o["sandbox"] = p
	}
}

// referrerpolicy 配置
func iFrame_referrerpolicy(p string) opt {
	return func(o map[string]interface{}) {
		o["referrerpolicy"] = p
	}
}

// iframe 高度
func iFrame_height(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// iframe 宽度
func iFrame_width(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}
