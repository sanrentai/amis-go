package amis

// 提示
func Alert(opts ...opt) map[string]interface{} {
	return newCompent("alert", opts...)
}

// 外层 Dom 的类名
func Alert_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 级别，可以是：`info`、`success`、`warning` 或者 `danger`
func Alert_level(p string) opt {
	return func(o map[string]interface{}) {
		o["level"] = p
	}
}

// 显示内容
func Alert_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 是否显示关闭按钮
func Alert_showCloseButton(p bool) opt {
	return func(o map[string]interface{}) {
		o["showCloseButton"] = p
	}
}

// 关闭按钮的 CSS 类名
func Alert_closeButtonClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["closeButtonClassName"] = p
	}
}

// 是否显示 icon
func Alert_showIcon(p bool) opt {
	return func(o map[string]interface{}) {
		o["showIcon"] = p
	}
}

// 自定义 icon
func Alert_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// icon 的 CSS 类名
func Alert_iconClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["iconClassName"] = p
	}
}
