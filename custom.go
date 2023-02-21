package amis

// 自定义组件
func Custom(opts ...opt) map[string]interface{} {
	return newCompent("custom", opts...)
}

// 节点 id
func Custom_id(p string) opt {
	return func(o map[string]interface{}) {
		o["id"] = p
	}
}

// 节点 名称
func Custom_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 节点 class
func Custom_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 默认使用 div 标签，如果 true 就使用 span 标签
func Custom_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
}

// 初始化节点 html
func Custom_html(p string) opt {
	return func(o map[string]interface{}) {
		o["html"] = p
	}
}

// 节点初始化之后调的用函数
func Custom_onMount(p string) opt {
	return func(o map[string]interface{}) {
		o["onMount"] = p
	}
}

// 数据有更新的时候调用的函数
func Custom_onUpdate(p string) opt {
	return func(o map[string]interface{}) {
		o["onUpdate"] = p
	}
}

// 节点销毁的时候调用的函数
func Custom_onUnmount(p string) opt {
	return func(o map[string]interface{}) {
		o["onUnmount"] = p
	}
}
