package amis

// 页面
func Page(opts ...opt) map[string]interface{} {
	return newCompent("page", opts...)
}

// 页面标题
func Page_title(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 页面副标题
func Page_subTitle(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["subTitle"] = p
	}
}

// 标题附近会出现一个提示图标，鼠标放上去会提示该内容。
func Page_remark(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["remark"] = p
	}
}

// 往页面的边栏区域加内容
func Page_aside(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["aside"] = p
	}
}

// 页面的边栏区域宽度是否可调整
func Page_asideResizor(p bool) opt {
	return func(o map[string]interface{}) {
		o["asideResizor"] = p
	}
}

// 页面边栏区域的最小宽度
func Page_asideMinWidth(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["asideMinWidth"] = p
	}
}

// 页面边栏区域的最大宽度
func Page_asideMaxWidth(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["asideMaxWidth"] = p
	}
}

// 用来控制边栏固定与否
func Page_asideSticky(p bool) opt {
	return func(o map[string]interface{}) {
		o["asideSticky"] = p
	}
}

// 往页面的右上角加内容，需要注意的是，当有 title 时，该区域在右上角，没有时该区域在顶部
func Page_toolbar(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["toolbar"] = p
	}
}

// 往页面的内容区域加内容
func Page_body(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}

// 外层 dom 类名
func Page_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 自定义 CSS 变量，请参考[样式](../style)
func Page_cssVars(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["cssVars"] = p
	}
}

// Toolbar dom 类名
func Page_toolbarClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["toolbarClassName"] = p
	}
}

// Body dom 类名
func Page_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
}

// Aside dom 类名
func Page_asideClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["asideClassName"] = p
	}
}

// Header 区域 dom 类名
func Page_headerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headerClassName"] = p
	}
}

// Page 用来获取初始数据的 api。返回的数据可以整个 page 级别使用。
func Page_initApi(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initApi"] = p
	}
}

// 是否起始拉取 initApi
func Page_initFetch(p bool) opt {
	return func(o map[string]interface{}) {
		o["initFetch"] = p
	}
}

// 是否起始拉取 initApi, 通过表达式配置
func Page_initFetchOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["initFetchOn"] = p
	}
}

// 刷新时间(最小 1000)
func Page_interval(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 配置刷新时是否显示加载动画
func Page_silentPolling(p bool) opt {
	return func(o map[string]interface{}) {
		o["silentPolling"] = p
	}
}

// 通过表达式来配置停止刷新的条件
func Page_stopAutoRefreshWhen(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["stopAutoRefreshWhen"] = p
	}
}

// 下拉刷新配置（仅用于移动端）
func Page_pullRefresh(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["pullRefresh"] = p
	}
}
