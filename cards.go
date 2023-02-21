package amis

// 卡片组
func Cards(opts ...opt) map[string]interface{} {
	return newCompent("cards", opts...)
}

// 标题
func Cards_title(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 数据源, 获取当前数据域中的变量
func Cards_source(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 当没数据的时候的文字提示
func Cards_placeholder(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 外层 CSS 类名
func Cards_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 顶部外层 CSS 类名
func Cards_headerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["headerClassName"] = p
	}
}

// 底部外层 CSS 类名
func Cards_footerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["footerClassName"] = p
	}
}

// 卡片 CSS 类名
func Cards_itemClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["itemClassName"] = p
	}
}

// 配置卡片信息
func Cards_card(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["card"] = p
	}
}
