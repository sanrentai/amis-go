package amis

// Component
func Web(opts ...opt) map[string]interface{} {
	return newCompent("web-component", opts...)
}

// 具体使用的 web-component 标签
func Web_tag(p string) opt {
	return func(o map[string]interface{}) {
		o["tag"] = p
	}
}

// 标签上的属性
func Web_props(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["props"] = p
	}
}

// 子节点
func Web_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}
