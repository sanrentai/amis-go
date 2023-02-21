package amis

// 按钮
func Button(opts ...opt) map[string]interface{} {
	return newCompent("button", opts...)
}

// 指定添加 button 类名
func Button_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
func Button_url(p string) opt {
	return func(o map[string]interface{}) {
		o["url"] = p
	}
}

// 按钮失效状态
func Button_disabled(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["disabled"] = p
	}
}

// 将按钮宽度调整为其父宽度的选项
func Button_block(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["block"] = p
	}
}

// 显示按钮 loading 效果
func Button_loading(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["loading"] = p
	}
}

// 显示按钮 loading 表达式
func Button_loadingOn(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["loadingOn"] = p
	}
}
