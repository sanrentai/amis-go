package amis

// 渲染
func Markdown(opts ...opt) map[string]interface{} {
	return newCompent("markdown", opts...)
}

// 名称
func Markdown_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 静态值
func Markdown_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 类名
func Markdown_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 外部地址
func Markdown_src(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}
