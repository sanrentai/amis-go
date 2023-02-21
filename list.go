package amis

// 列表
func List(opts ...opt) map[string]interface{} {
	return newCompent("list", opts...)
}

// 标题
func List_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 数据源, 获取当前数据域变量，支持[数据映射](../../docs/concepts/data-mapping)
func List_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 当没数据的时候的文字提示
func List_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 外层 CSS 类名
func List_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 顶部外层 CSS 类名
func List_headerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headerClassName"] = p
	}
}

// 底部外层 CSS 类名
func List_footerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["footerClassName"] = p
	}
}

// 配置单条信息
func List_listItem(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["listItem"] = p
	}
}

// // 标题
// func List_listItem.title(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.title"] = p
// 	}
// }

// // 标题 CSS 类名
// func List_listItem.titleClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.titleClassName"] = p
// 	}
// }

// // 副标题
// func List_listItem.subTitle(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.subTitle"] = p
// 	}
// }

// // 图片地址
// func List_listItem.avatar(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.avatar"] = p
// 	}
// }

// // 图片 CSS 类名
// func List_listItem.avatarClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.avatarClassName"] = p
// 	}
// }

// // 描述
// func List_listItem.desc(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.desc"] = p
// 	}
// }

// // 内容容器，主要用来放置非表单项组件
// func List_listItem.body(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.body"] = p
// 	}
// }

// // 按钮区域
// func List_listItem.actions(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.actions"] = p
// 	}
// }

// // 按钮位置
// func List_listItem.actionsPosition(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["listItem.actionsPosition"] = p
// 	}
// }

// 版本
func List_事件名称(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["事件名称"] = p
	}
}

// `2.4.0`
func List_itemClick(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemClick"] = p
	}
}

// 当前行所在数据域
func List_data(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["data"] = p
	}
}

// 行索引值，从 0 开始
func List_index(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["index"] = p
	}
}
