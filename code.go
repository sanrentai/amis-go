package amis

// 代码高亮
func Code(opts ...opt) map[string]interface{} {
	return newCompent("code", opts...)
}

// 外层 CSS 类名
func Code_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 显示的颜色值
func Code_value(p string) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 在其他组件中，时，用作变量映射
func Code_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 所使用的高亮语言，默认是 plaintext
func Code_language(p string) opt {
	return func(o map[string]interface{}) {
		o["language"] = p
	}
}

// 默认 tab 大小
func Code_tabSize(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["tabSize"] = p
	}
}

// 主题，还有 'vs-dark'
func Code_editorTheme(p string) opt {
	return func(o map[string]interface{}) {
		o["editorTheme"] = p
	}
}

// 是否折行
func Code_wordWrap(p string) opt {
	return func(o map[string]interface{}) {
		o["wordWrap"] = p
	}
}
