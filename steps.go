package amis

// 步骤条
func Steps(opts ...opt) map[string]interface{} {
	return newCompent("steps", opts...)
}

// 数组，配置步骤信息
func Steps_steps(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["steps"] = p
	}
}

// 选项组源，可通过数据映射获取当前数据域变量、或者配置 API 对象
func Steps_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 关联上下文变量
func Steps_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 自定义类名
func Steps_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 点状步骤条
func Steps_progressDot(p bool) opt {
	return func(o map[string]interface{}) {
		o["progressDot"] = p
	}
}

// 说明
func Steps_属性名(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["属性名"] = p
	}
}

// icon 名，支持 fontawesome v4 或使用 url
func Steps_icon(p string) opt {
	return func(o map[string]interface{}) {
		o["icon"] = p
	}
}

// value
func Steps_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}
