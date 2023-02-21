package amis

// 面板
func Panel(opts ...opt) map[string]interface{} {
	return newCompent("panel", opts...)
}

// 外层 Dom 的类名
func Panel_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// header 区域的类名
func Panel_headerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headerClassName"] = p
	}
}

// footer 区域的类名
func Panel_footerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["footerClassName"] = p
	}
}

// actions 区域的类名
func Panel_actionsClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["actionsClassName"] = p
	}
}

// body 区域的类名
func Panel_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
}

// 标题
func Panel_title(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 头部容器
func Panel_header(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["header"] = p
	}
}

// 内容容器
func Panel_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 底部容器
func Panel_footer(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["footer"] = p
	}
}

// 是否固定底部容器
func Panel_affixFooter(p bool) opt {
	return func(o map[string]interface{}) {
		o["affixFooter"] = p
	}
}

// 按钮区域
func Panel_actions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["actions"] = p
	}
}
