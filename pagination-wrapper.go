package amis

// 分页容器
func PaginationWrapper(opts ...opt) map[string]interface{} {
	return newCompent("pagination-wrapper", opts...)
}

// 是否显示快速跳转输入框
func PaginationWrapper_showPageInput(p bool) opt {
	return func(o map[string]interface{}) {
		o["showPageInput"] = p
	}
}

// 最多显示多少个分页按钮
func PaginationWrapper_maxButtons(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxButtons"] = p
	}
}

// 输入字段名
func PaginationWrapper_inputName(p string) opt {
	return func(o map[string]interface{}) {
		o["inputName"] = p
	}
}

// 输出字段名
func PaginationWrapper_outputName(p string) opt {
	return func(o map[string]interface{}) {
		o["outputName"] = p
	}
}

// 每页显示多条数据
func PaginationWrapper_perPage(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["perPage"] = p
	}
}

// 分页显示位置，如果配置为 none 则需要自己在内容区域配置 pagination 组件，否则不显示
func PaginationWrapper_position(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["position"] = p
	}
}

// 内容区域
func PaginationWrapper_body(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["body"] = p
	}
}
