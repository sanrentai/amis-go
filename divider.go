package amis

// 分割线
func Divider(opts ...opt) map[string]interface{} {
	return newCompent("divider", opts...)
}

// 外层 Dom 的类名
func Divider_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 分割线的样式，支持`dashed`和`solid`
func Divider_lineStyle(p string) opt {
	return func(o map[string]interface{}) {
		o["lineStyle"] = p
	}
}
