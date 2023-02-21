package amis

// 标签
func Tag(opts ...opt) map[string]interface{} {
	return newCompent("tag", opts...)
}

// 标签内容
func Tag_label(p string) opt {
	return func(o map[string]interface{}) {
		o["label"] = p
	}
}

// status 模式下的前置图标
func Tag_icon(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 自定义 CSS 样式类名
func Tag_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 自定义样式（行内样式），优先级最高
func Tag_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}
