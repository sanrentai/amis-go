package amis

// 包裹容器
func Wrapper(opts ...opt) map[string]interface{} {
	return newCompent("wrapper", opts...)
}

// 外层 Dom 的类名
func Wrapper_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 支持: `xs`、`sm`、`md`和`lg`
func Wrapper_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
}

// 内容容器
func Wrapper_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}
