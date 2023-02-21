package amis

// 容器
func Container(opts ...opt) map[string]interface{} {
	return newCompent("container", opts...)
}

// 外层 Dom 的类名
func Container_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 容器内容区的类名
func Container_bodyClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["bodyClassName"] = p
	}
}

// 容器标签名
func Container_wrapperComponent(p string) opt {
	return func(o map[string]interface{}) {
		o["wrapperComponent"] = p
	}
}

// 自定义样式
func Container_style(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["style"] = p
	}
}

// 容器内容
func Container_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}
