package amis

// 导航
func Nav(opts ...opt) map[string]interface{} {
	return newCompent("nav", opts...)
}

// 外层 Dom 的类名
func Nav_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 设置成 false 可以以 tabs 的形式展示
func Nav_stacked(p bool) opt {
	return func(o map[string]interface{}) {
		o["stacked"] = p
	}
}

// 可以通过变量或 API 接口动态创建导航
func Nav_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 用来延时加载选项详情的接口，可以不配置，不配置公用 source 接口。
func Nav_deferApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["deferApi"] = p
	}
}

// 更多操作相关配置
func Nav_itemActions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemActions"] = p
	}
}

// 是否支持拖拽排序
func Nav_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
}

// 仅允许同层级内拖拽
func Nav_dragOnSameLevel(p bool) opt {
	return func(o map[string]interface{}) {
		o["dragOnSameLevel"] = p
	}
}

// 保存排序的 api
func Nav_saveOrderApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["saveOrderApi"] = p
	}
}

// 角标
func Nav_itemBadge(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemBadge"] = p
	}
}

// 链接集合
func Nav_links(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["links"] = p
	}
}

// // 名称
// func Nav_links[x].label(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].label"] = p
// 	}
// }

// // 链接地址
// func Nav_links[x].to(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].to"] = p
// 	}
// }

// //
// func Nav_links[x].target(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].target"] = p
// 	}
// }

// // 图标
// func Nav_links[x].icon(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].icon"] = p
// 	}
// }

// // 子链接
// func Nav_links[x].children(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].children"] = p
// 	}
// }

// // 初始是否展开
// func Nav_links[x].unfolded(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].unfolded"] = p
// 	}
// }

// // 是否高亮
// func Nav_links[x].active(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].active"] = p
// 	}
// }

// // 是否高亮的条件，留空将自动分析链接地址
// func Nav_links[x].activeOn(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].activeOn"] = p
// 	}
// }

// // 标记是否为懒加载项
// func Nav_links[x].defer(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].defer"] = p
// 	}
// }

// // 可以不配置，如果配置优先级更高
// func Nav_links[x].deferApi(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].deferApi"] = p
// 	}
// }

// 响应式收纳配置
func Nav_overflow(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["overflow"] = p
	}
}

// // 是否开启响应式收纳
// func Nav_overflow.enable(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.enable"] = p
// 	}
// }

// // 菜单触发按钮的图标
// func Nav_overflow.overflowIndicator(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.overflowIndicator"] = p
// 	}
// }

// // 开启响应式收纳后导航最大可显示数量，超出此数量的导航将被收纳到下拉菜单中，默认为自动计算
// func Nav_overflow.maxVisibleCount(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.maxVisibleCount"] = p
// 	}
// }

// // 包裹导航的外层标签名，可以使用其他标签渲染
// func Nav_overflow.wrapperComponent(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.wrapperComponent"] = p
// 	}
// }

// // 自定义样式
// func Nav_overflow.style(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.style"] = p
// 	}
// }

// // 菜单按钮 CSS 类名
// func Nav_overflow.overflowClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.overflowClassName"] = p
// 	}
// }

// // Popover 浮层 CSS 类名
// func Nav_overflow.overflowPopoverClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.overflowPopoverClassName"] = p
// 	}
// }

// // 菜单外层 CSS 类名
// func Nav_overflow.overflowListClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["overflow.overflowListClassName"] = p
// 	}
// }
