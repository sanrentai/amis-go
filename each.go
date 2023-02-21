package amis

// 循环渲染器
func Each(opts ...opt) map[string]interface{} {
	return newCompent("each", opts...)
}

// 用于循环的值
func Each_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 获取数据域中变量
func Each_name(p string) opt {
	return func(o map[string]interface{}) {
		o["name"] = p
	}
}

// 获取数据域中变量， 支持 [数据映射](../../docs/concepts/data-mapping)
func Each_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 使用`value`中的数据，循环输出渲染器。
func Each_items(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["items"] = p
	}
}

// 当 `value` 值不存在或为空数组时的占位文本
func Each_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}
