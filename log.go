package amis

// 实时日志
func Log(opts ...opt) map[string]interface{} {
	return newCompent("log", opts...)
}

// 展示区域高度
func Log_height(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// 外层 CSS 类名
func Log_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 是否自动滚动
func Log_autoScroll(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoScroll"] = p
	}
}

// 加载中的文字
func Log_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 返回内容的字符编码
func Log_encoding(p string) opt {
	return func(o map[string]interface{}) {
		o["encoding"] = p
	}
}

// 接口
func Log_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 设置每行高度，将会开启虚拟渲染
func Log_rowHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rowHeight"] = p
	}
}

// 最大显示行数
func Log_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
}

// 可选日志操作：['stop','restart',clear','showLineNumber','filter']
func Log_operation(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["operation"] = p
	}
}
