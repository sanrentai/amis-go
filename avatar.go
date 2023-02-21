package amis

// 头像
func Avatar(opts ...opt) map[string]interface{} {
	return newCompent("avatar", opts...)
}

// 外层 dom 的类名
func Avatar_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 外层 dom 的样式
func Avatar_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}

// 图片地址
func Avatar_src(p string) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}

// 文字
func Avatar_text(p string) opt {
	return func(o map[string]interface{}) {
		o["text"] = p
	}
}

// 图标
func Avatar_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// 控制字符类型距离左右两侧边界单位像素
func Avatar_gap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["gap"] = p
	}
}

// 图像无法显示时的替代文本
func Avatar_alt(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["alt"] = p
	}
}

// 图片是否允许拖动
func Avatar_draggable(p bool) opt {
	return func(o map[string]interface{}) {
		o["draggable"] = p
	}
}

// 图片加载失败的字符串，这个字符串是一个New Function内部执行的字符串，参数是event（使用event.nativeEvent获取原生dom事件），这个字符串需要返回boolean值。设置 `"return ture;"` 会在图片加载失败后，使用 `text` 或者 `icon` 代表的信息来进行替换。目前图片加载失败默认是不进行置换。注意：图片加载失败，不包括$获取数据为空情况
func Avatar_onError(p string) opt {
	return func(o map[string]interface{}) {
		o["onError"] = p
	}
}
