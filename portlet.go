package amis

// 门户栏目
func Portlet(opts ...opt) map[string]interface{} {
	return newCompent("portlet", opts...)
}

// 外层 Dom 的类名
func Portlet_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// Tabs Dom 的类名
func Portlet_tabsClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["tabsClassName"] = p
	}
}

// Tabs content Dom 的类名
func Portlet_contentClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["contentClassName"] = p
	}
}

// tabs 内容
func Portlet_tabs(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tabs"] = p
	}
}

// tabs 关联数据，关联后可以重复生成选项卡
func Portlet_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// tabs 中的工具栏，不随 tab 切换而变化
func Portlet_toolbar(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["toolbar"] = p
	}
}

// 标题右侧信息
func Portlet_description(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["description"] = p
	}
}

// 隐藏头部
func Portlet_hideHeader(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideHeader"] = p
	}
}

// 去掉分隔线
func Portlet_divider(p bool) opt {
	return func(o map[string]interface{}) {
		o["divider"] = p
	}
}

// Tab 标题
// func Portlet_tabs[x].title(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].title"] = p
// 	}
// }

// // Tab 的图标
// func Portlet_tabs[x].icon(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].icon"] = p
// 	}
// }

// // 内容区
// func Portlet_tabs[x].tab(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].tab"] = p
// 	}
// }

// // tabs 中的工具栏，随 tab 切换而变化
// func Portlet_tabs[x].toolbar(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].toolbar"] = p
// 	}
// }

// // 设置以后内容每次都会重新渲染，对于 crud 的重新拉取很有用
// func Portlet_tabs[x].reload(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].reload"] = p
// 	}
// }

// // 每次退出都会销毁当前 tab 栏内容
// func Portlet_tabs[x].unmountOnExit(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].unmountOnExit"] = p
// 	}
// }

// // Tab 区域样式
// func Portlet_tabs[x].className(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].className"] = p
// 	}
// }

// 只有在点中 tab 的时候才渲染
func Portlet_mountOnEnter(p bool) opt {
	return func(o map[string]interface{}) {
		o["mountOnEnter"] = p
	}
}

// 切换 tab 的时候销毁
func Portlet_unmountOnExit(p bool) opt {
	return func(o map[string]interface{}) {
		o["unmountOnExit"] = p
	}
}

// 是否导航支持内容溢出滚动，`vertical`和`chrome`模式下不支持该属性；`chrome`模式默认压缩标签
func Portlet_scrollable(p bool) opt {
	return func(o map[string]interface{}) {
		o["scrollable"] = p
	}
}
