package amis

// 图标
func Icon(opts ...opt) map[string]interface{} {
	return newCompent("icon", opts...)
}

// 外层 CSS 类名
func Icon_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// icon 名称，支持 [fontawesome v4](https://fontawesome.com/v4/icons/) 或 通过 registerIcon 注册的 icon、或使用 url
func Icon_icon(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// icon 类型，默认为`fa`, 表示 fontawesome v4。也支持 iconfont, 如果是 fontawesome v5 以上版本或者其他框架可以设置为空字符串
func Icon_vendor(p string) opt {
	return func(o map[string]interface{}) {
		o["vendor"] = p
	}
}
