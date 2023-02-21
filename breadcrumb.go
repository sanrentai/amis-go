package amis

// 面包屑
func Breadcrumb(opts ...opt) map[string]interface{} {
	return newCompent("breadcrumb", opts...)
}

// 外层类名
func Breadcrumb_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 导航项类名
func Breadcrumb_itemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemClassName"] = p
	}
}

// 分割符类名
func Breadcrumb_separatorClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["separatorClassName"] = p
	}
}

// 下拉菜单类名
func Breadcrumb_dropdownClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["dropdownClassName"] = p
	}
}

// 下拉菜单项类名
func Breadcrumb_dropdownItemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["dropdownItemClassName"] = p
	}
}

// 分隔符
func Breadcrumb_separator(p string) opt {
	return func(o map[string]interface{}) {
		o["separator"] = p
	}
}

// 最大展示长度
func Breadcrumb_labelMaxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["labelMaxLength"] = p
	}
}

// 动态数据
func Breadcrumb_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

func Breadcrumb_items(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
}

// // 文本
// func Breadcrumb_items[].label(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].label"] = p
// 	}
// }

// // 链接
// func Breadcrumb_items[].href(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].href"] = p
// 	}
// }

// // [图标](icon)
// func Breadcrumb_items[].icon(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].icon"] = p
// 	}
// }

// // 下拉菜单，dropdown[]的每个对象都包含label、href、icon属性
// func Breadcrumb_items[].dropdown(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].dropdown"] = p
// 	}
// }
