package amis

// 锚点导航
func AnchorNav(opts ...opt) map[string]interface{} {
	return newCompent("anchor-nav", opts...)
}

// 外层 Dom 的类名
func AnchorNav_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 导航 Dom 的类名
func AnchorNav_linkClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["linkClassName"] = p
	}
}

// 锚点区域 Dom 的类名
func AnchorNav_sectionClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["sectionClassName"] = p
	}
}

// links 内容
func AnchorNav_links(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["links"] = p
	}
}

// 可以配置导航水平展示还是垂直展示。对应的配置项分别是：vertical、horizontal
func AnchorNav_direction(p string) opt {
	return func(o map[string]interface{}) {
		o["direction"] = p
	}
}

// 需要定位的区域
func AnchorNav_active(p string) opt {
	return func(o map[string]interface{}) {
		o["active"] = p
	}
}

// // 区域 标题
// func AnchorNav_links[x].title(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].title"] = p
// 	}
// }

// // 区域 标识
// func AnchorNav_links[x].href(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].href"] = p
// 	}
// }

// // 区域 内容区
// func AnchorNav_links[x].body(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].body"] = p
// 	}
// }

// // 区域成员 样式
// func AnchorNav_links[x].className(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["links[x].className"] = p
// 	}
// }
