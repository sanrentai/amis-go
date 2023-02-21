package amis

// 布局
func Flex(opts ...opt) map[string]interface{} {
	return newCompent("flex", opts...)
}

// css 类名
func Flex_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// "start", "flex-start", "center", "end", "flex-end", "space-around", "space-between", "space-evenly"
func Flex_justify(p string) opt {
	return func(o map[string]interface{}) {
		o["justify"] = p
	}
}

// "stretch", "start", "flex-start", "flex-end", "end", "center", "baseline"
func Flex_alignItems(p string) opt {
	return func(o map[string]interface{}) {
		o["alignItems"] = p
	}
}

// 自定义样式
func Flex_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}
