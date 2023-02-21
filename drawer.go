package amis

// 抽屉
func Drawer(opts ...opt) map[string]interface{} {
	return newCompent("drawer", opts...)
}

// 弹出层标题
func Drawer_title(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 往 Drawer 内容区加内容
func Drawer_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 指定 Drawer 大小，支持: `xs`、`sm`、`md`、`lg`、`xl`
func Drawer_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// 指定 Drawer 方向，支持: `left`、`right`、`top`、`bottom`
func Drawer_position(p string) opt {
	return func(o map[string]interface{}) {
		o["position"] = p
	}
}

// Drawer 最外层容器的样式类名
func Drawer_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// Drawer 头部 区域的样式类名
func Drawer_headerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headerClassName"] = p
	}
}

// Drawer body 区域的样式类名
func Drawer_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
}

// Drawer 页脚 区域的样式类名
func Drawer_footerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["footerClassName"] = p
	}
}

// 是否展示关闭按钮，当值为 false 时，默认开启 closeOnOutside
func Drawer_showCloseButton(p bool) opt {
	return func(o map[string]interface{}) {
		o["showCloseButton"] = p
	}
}

// 是否支持按 `Esc` 关闭 Drawer
func Drawer_closeOnEsc(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnEsc"] = p
	}
}

// 点击内容区外是否关闭 Drawer
func Drawer_closeOnOutside(p bool) opt {
	return func(o map[string]interface{}) {
		o["closeOnOutside"] = p
	}
}

// 是否显示蒙层
func Drawer_overlay(p bool) opt {
	return func(o map[string]interface{}) {
		o["overlay"] = p
	}
}

// 是否可通过拖拽改变 Drawer 大小
func Drawer_resizable(p bool) opt {
	return func(o map[string]interface{}) {
		o["resizable"] = p
	}
}

// 可以不设置，默认只有两个按钮。
func Drawer_actions(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["actions"] = p
	}
}

// 支持 [数据映射](../../docs/concepts/data-mapping)，如果不设定将默认将触发按钮的上下文中继承数据。
func Drawer_data(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["data"] = p
	}
}
