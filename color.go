package amis

// 颜色
func Color(opts ...opt) map[string]interface{} {
	return newCompent("color", opts...)
}

// 外层 CSS 类名
func Color_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 显示的颜色值
func Color_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 在其他组件中，时，用作变量映射
func Color_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 默认颜色值
func Color_defaultColor(p string) opt {
	return func(o map[string]interface{}) {
		o["defaultColor"] = p
	}
}

// 是否显示右边的颜色值
func Color_showValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["showValue"] = p
	}
}
