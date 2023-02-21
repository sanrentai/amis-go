package amis

// 图片
func Image(opts ...opt) map[string]interface{} {
	return newCompent("image", opts...)
}

// 外层 CSS 类名
func Image_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 组件内层 CSS 类名
func Image_innerClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["innerClassName"] = p
	}
}

// 图片 CSS 类名
func Image_imageClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["imageClassName"] = p
	}
}

// 图片缩率图 CSS 类名
func Image_thumbClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["thumbClassName"] = p
	}
}

// 图片缩率高度
func Image_height(p string) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// 图片缩率宽度
func Image_width(p string) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}

// 标题
func Image_title(p string) opt {
	return func(o map[string]interface{}) {
		o["title"] = p
	}
}

// 描述
func Image_imageCaption(p string) opt {
	return func(o map[string]interface{}) {
		o["imageCaption"] = p
	}
}

// 占位文本
func Image_placeholder(p string) opt {
	return func(o map[string]interface{}) {
		o["placeholder"] = p
	}
}

// 无数据时显示的图片
func Image_defaultImage(p string) opt {
	return func(o map[string]interface{}) {
		o["defaultImage"] = p
	}
}

// 缩略图地址
func Image_src(p string) opt {
	return func(o map[string]interface{}) {
		o["src"] = p
	}
}

// 外部链接地址
func Image_href(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["href"] = p
	}
}

// 原图地址
func Image_originalSrc(p string) opt {
	return func(o map[string]interface{}) {
		o["originalSrc"] = p
	}
}

// 支持放大预览
func Image_enlargeAble(p bool) opt {
	return func(o map[string]interface{}) {
		o["enlargeAble"] = p
	}
}

// 放大预览的标题
func Image_enlargeTitle(p string) opt {
	return func(o map[string]interface{}) {
		o["enlargeTitle"] = p
	}
}

// 放大预览的描述
func Image_enlargeCaption(p string) opt {
	return func(o map[string]interface{}) {
		o["enlargeCaption"] = p
	}
}

// 预览图模式，可选：`'w-full'`, `'h-full'`, `'contain'`, `'cover'`
func Image_thumbMode(p string) opt {
	return func(o map[string]interface{}) {
		o["thumbMode"] = p
	}
}

// 预览图比例，可选：`'1:1'`, `'4:3'`, `'16:9'`
func Image_thumbRatio(p string) opt {
	return func(o map[string]interface{}) {
		o["thumbRatio"] = p
	}
}

// 图片展示模式，可选：`'thumb'`, `'original'` 即：缩略图模式 或者 原图模式
func Image_imageMode(p string) opt {
	return func(o map[string]interface{}) {
		o["imageMode"] = p
	}
}
