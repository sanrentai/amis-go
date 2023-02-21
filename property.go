package amis

// 属性表
func Property(opts ...opt) map[string]interface{} {
	return newCompent("property", opts...)
}

// 外层 dom 的类名
func Property_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 外层 dom 的样式
func Property_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}

// 属性名的样式
func Property_labelStyle(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["labelStyle"] = p
	}
}

// 属性值的样式
func Property_contentStyle(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["contentStyle"] = p
	}
}

// 每行几列
func Property_column(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["column"] = p
	}
}

// 显示模式，目前只有 'table' 和 'simple'
func Property_mode(p string) opt {
	return func(o map[string]interface{}) {
		o["mode"] = p
	}
}

// 'simple' 模式下属性名和值之间的分隔符
func Property_separator(p string) opt {
	return func(o map[string]interface{}) {
		o["separator"] = p
	}
}

// 标题
func Property_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 数据源
func Property_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 属性名
// func Property_items[].label(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].label"] = p
// 	}
// }

// // 属性值
// func Property_items[].content(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].content"] = p
// 	}
// }

// // 属性值跨几列
// func Property_items[].span(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].span"] = p
// 	}
// }

// // 显示表达式
// func Property_items[].visibleOn(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].visibleOn"] = p
// 	}
// }

// // 隐藏表达式
// func Property_items[].hiddenOn(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["items[].hiddenOn"] = p
// 	}
// }
