package amis

// 链接
func Link(opts ...opt) map[string]interface{} {
	return newCompent("link", opts...)
}

// 标签内文本
func Link_body(p string) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 链接地址
func Link_href(p string) opt {
	return func(o map[string]interface{}) {
		o["href"] = p
	}
}

// 是否在新标签页打开
func Link_blank(p bool) opt {
	return func(o map[string]interface{}) {
		o["blank"] = p
	}
}

// a 标签的 target，优先于 blank 属性
func Link_htmlTarget(p string) opt {
	return func(o map[string]interface{}) {
		o["htmlTarget"] = p
	}
}

// a 标签的 title
func Link_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 禁用超链接
func Link_disabled(p bool) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
}

// 超链接图标，以加强显示
func Link_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 右侧图标
func Link_rightIcon(p string) opt {
	return func(o map[string]interface{}) {
		o["rightIcon"] = p
	}
}
