package amis

// 模板
func Tpl(opts ...opt) map[string]interface{} {
	return newCompent("tpl", opts...)
}

// 外层 Dom 的类名
func Tpl_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 配置模板
func Tpl_tpl(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tpl"] = p
	}
}

// 是否设置外层 DOM 节点的 title 属性为文本内容
func Tpl_showNativeTitle(p bool) opt {
	return func(o map[string]interface{}) {
		o["showNativeTitle"] = p
	}
}

func Tpl_inline(p bool) opt {
	return func(o map[string]interface{}) {
		o["inline"] = p
	}
}
