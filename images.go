package amis

// 图片集
func Images(opts ...opt) map[string]interface{} {
	return newCompent("images", opts...)
}

// 外层 CSS 类名
func Images_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 默认展示图片
func Images_defaultImage(p string) opt {
	return func(o map[string]interface{}) {
		o["defaultImage"] = p
	}
}

// 图片数组
func Images_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// 数据源
func Images_source(p string) opt {
	return func(o map[string]interface{}) {
		o["source"] = p
	}
}

// 分隔符，当 value 为字符串时，用该值进行分隔拆分
func Images_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
}

// 预览图地址，支持数据映射获取对象中图片变量
func Images_src(p string) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}

// 原图地址，支持数据映射获取对象中图片变量
func Images_originalSrc(p string) opt {
	return func(o map[string]interface{}) {
		o["originalSrc"] = p
	}
}

// 支持放大预览
func Images_enlargeAble(p bool) opt {
	return func(o map[string]interface{}) {
		o["enlargeAble"] = p
	}
}

// 预览图模式，可选：`'w-full'`, `'h-full'`, `'contain'`, `'cover'`
func Images_thumbMode(p string) opt {
	return func(o map[string]interface{}) {
		o["thumbMode"] = p
	}
}

// 预览图比例，可选：`'1:1'`, `'4:3'`, `'16:9'`
func Images_thumbRatio(p string) opt {
	return func(o map[string]interface{}) {
		o["thumbRatio"] = p
	}
}
