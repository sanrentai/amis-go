package amis

// 选项卡
func Tabs(opts ...opt) map[string]interface{} {
	return newCompent("tabs", opts...)
}

// 组件初始化时激活的选项卡，hash 值或索引值，支持使用表达式 `2.7.1 以上版本`
func Tabs_defaultKey(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["defaultKey"] = p
	}
}

// 激活的选项卡，hash 值或索引值，支持使用表达式，可响应上下文数据变化
func Tabs_activeKey(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["activeKey"] = p
	}
}

// 外层 Dom 的类名
func Tabs_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 展示模式，取值可以是 `line`、`card`、`radio`、`vertical`、`chrome`、`simple`、`strong`、`tiled`、`sidebar`
func Tabs_tabsMode(p string) opt {
	return func(o map[string]interface{}) {
		o["tabsMode"] = p
	}
}

// Tabs Dom 的类名
func Tabs_tabsClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["tabsClassName"] = p
	}
}

// tabs 内容
func Tabs_tabs(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tabs"] = p
	}
}

// tabs 关联数据，关联后可以重复生成选项卡
func Tabs_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// tabs 中的工具栏
func Tabs_toolbar(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["toolbar"] = p
	}
}

// tabs 中工具栏的类名
func Tabs_toolbarClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["toolbarClassName"] = p
	}
}

// // Tab 标题
// func Tabs_tabs[x].title(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].title"] = p
// 	}
// }

// // Tab 的图标
// func Tabs_tabs[x].icon(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].icon"] = p
// 	}
// }

// // Tab 的图标位置
// func Tabs_tabs[x].iconPosition(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].iconPosition"] = p
// 	}
// }

// // 内容区
// func Tabs_tabs[x].tab(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].tab"] = p
// 	}
// }

// // 设置以后将跟 url 的 hash 对应
// func Tabs_tabs[x].hash(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].hash"] = p
// 	}
// }

// // 设置以后内容每次都会重新渲染，对于 crud 的重新拉取很有用
// func Tabs_tabs[x].reload(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].reload"] = p
// 	}
// }

// // 每次退出都会销毁当前 tab 栏内容
// func Tabs_tabs[x].unmountOnExit(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].unmountOnExit"] = p
// 	}
// }

// // Tab 区域样式
// func Tabs_tabs[x].className(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].className"] = p
// 	}
// }

// // 是否支持删除，优先级高于组件的 `closable`
// func Tabs_tabs[x].closable(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].closable"] = p
// 	}
// }

// // 是否禁用
// func Tabs_tabs[x].disabled(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["tabs[x].disabled"] = p
// 	}
// }

// 只有在点中 tab 的时候才渲染
func Tabs_mountOnEnter(p bool) opt {
	return func(o map[string]interface{}) {
		o["mountOnEnter"] = p
	}
}

// 切换 tab 的时候销毁
func Tabs_unmountOnExit(p bool) opt {
	return func(o map[string]interface{}) {
		o["unmountOnExit"] = p
	}
}

// 是否支持新增
func Tabs_addable(p bool) opt {
	return func(o map[string]interface{}) {
		o["addable"] = p
	}
}

// 新增按钮文案
func Tabs_addBtnText(p string) opt {
	return func(o map[string]interface{}) {
		o["addBtnText"] = p
	}
}

// 是否支持删除
func Tabs_closable(p bool) opt {
	return func(o map[string]interface{}) {
		o["closable"] = p
	}
}

// 是否支持拖拽
func Tabs_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
}

// 是否支持提示
func Tabs_showTip(p bool) opt {
	return func(o map[string]interface{}) {
		o["showTip"] = p
	}
}

// 提示的类
func Tabs_showTipClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["showTipClassName"] = p
	}
}

// 收否可编辑标签名
func Tabs_editable(p bool) opt {
	return func(o map[string]interface{}) {
		o["editable"] = p
	}
}

// 是否导航支持内容溢出滚动。（属性废弃）
func Tabs_scrollable(p bool) opt {
	return func(o map[string]interface{}) {
		o["scrollable"] = p
	}
}

// `sidebar` 模式下，标签栏位置
func Tabs_sidePosition(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["sidePosition"] = p
	}
}

// 当 tabs 超出多少个时开始折叠
func Tabs_collapseOnExceed(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["collapseOnExceed"] = p
	}
}

// 用来设置折叠按钮的文字
func Tabs_collapseBtnLabel(p string) opt {
	return func(o map[string]interface{}) {
		o["collapseBtnLabel"] = p
	}
}

// 切换选项卡时触发
func Tabs_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
}

// 激活指定的选项卡
func Tabs_changeActiveKey(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["changeActiveKey"] = p
	}
}
