package amis

// 宫格导航
func GridNav(opts ...opt) map[string]interface{} {
	return newCompent("grid-nav", opts...)
}

// 外层 CSS 类名
func GridNav_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 列表项 css 类名
func GridNav_itemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemClassName"] = p
	}
}

// 图片数组
func GridNav_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 数据源
func GridNav_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 是否将列表项固定为正方形
func GridNav_square(p bool) opt {
	return func(o map[string]interface{}) {
		o["square"] = p
	}
}

// 是否将列表项内容居中显示
func GridNav_center(p bool) opt {
	return func(o map[string]interface{}) {
		o["center"] = p
	}
}

// 是否显示列表项边框
func GridNav_border(p bool) opt {
	return func(o map[string]interface{}) {
		o["border"] = p
	}
}

// 列表项之间的间距，默认单位为`px`
func GridNav_gutter(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["gutter"] = p
	}
}

// 是否调换图标和文本的位置
func GridNav_reverse(p bool) opt {
	return func(o map[string]interface{}) {
		o["reverse"] = p
	}
}

// 图标宽度占比，单位%
func GridNav_iconRatio(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["iconRatio"] = p
	}
}

// 列表项内容排列的方向，可选值为 `horizontal` 、`vertical`
func GridNav_direction(p string) opt {
	return func(o map[string]interface{}) {
		o["direction"] = p
	}
}

// 列数
func GridNav_columnNum(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["columnNum"] = p
	}
}

func GridNav_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
}

// // 列表项图标
// func GridNav_options.icon(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.icon"] = p
// 	}
// }

// // 列表项文案
// func GridNav_options.text(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.text"] = p
// 	}
// }

// // 列表项角标，详见 [Badge](./badge)
// func GridNav_options.badge(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.badge"] = p
// 	}
// }

// // 内部页面路径或外部跳转 URL 地址，优先级高于 clickAction
// func GridNav_options.link(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.link"] = p
// 	}
// }

// // 是否新页面打开，link 为 url 时有效
// func GridNav_options.blank(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.blank"] = p
// 	}
// }

// // 列表项点击交互 详见 [Action](./action)
// func GridNav_options.clickAction(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.clickAction"] = p
// 	}
// }
