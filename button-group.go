package amis

// 按钮组
func ButtonGroup(opts ...opt) map[string]interface{} {
	return newCompent("button-group", opts...)
}

// 是否使用垂直模式
func ButtonGroup_vertical(p bool) opt {
	return func(o map[string]interface{}) {
		o["vertical"] = p
	}
}

// 是否使用平铺模式
func ButtonGroup_tiled(p bool) opt {
	return func(o map[string]interface{}) {
		o["tiled"] = p
	}
}

// [按钮](./action)
func ButtonGroup_buttons(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["buttons"] = p
	}
}

// 外层 Dom 的类名
func ButtonGroup_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}
