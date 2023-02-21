package amis

// 角标
func Badge(opts ...opt) map[string]interface{} {
	return newCompent("badge", opts...)
}

// 角标类型，可以是 dot/text/ribbon
func Badge_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
}

// 角标大小
func Badge_size(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// 角标级别, 可以是 info/success/warning/danger, 设置之后角标背景颜色不同
func Badge_level(p string) opt {
	return func(o map[string]interface{}) {
		o["level"] = p
	}
}

// 设置封顶的数字值
func Badge_overflowCount(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["overflowCount"] = p
	}
}

// 角标位置， 可以是 top-right/top-left/bottom-right/bottom-left
func Badge_position(p string) opt {
	return func(o map[string]interface{}) {
		o["position"] = p
	}
}

// 角标位置，offset 相对于 position 位置进行水平、垂直偏移
func Badge_offset(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["offset"] = p
	}
}

// 外层 dom 的类名
func Badge_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 角标是否显示动画
func Badge_animation(p bool) opt {
	return func(o map[string]interface{}) {
		o["animation"] = p
	}
}

// 角标的自定义样式
func Badge_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}

// 控制角标的显示隐藏
func Badge_visibleOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["visibleOn"] = p
	}
}
