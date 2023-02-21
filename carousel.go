package amis

// 轮播图
func Carousel(opts ...opt) map[string]interface{} {
	return newCompent("carousel", opts...)
}

// 外层 Dom 的类名
func Carousel_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 轮播面板数据
func Carousel_options(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
}

// // 图片链接
// func Carousel_options.image(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.image"] = p
// 	}
// }

// // 图片打开网址的链接
// func Carousel_options.href(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.href"] = p
// 	}
// }

// // 图片类名
// func Carousel_options.imageClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.imageClassName"] = p
// 	}
// }

// // 图片标题
// func Carousel_options.title(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.title"] = p
// 	}
// }

// // 图片标题类名
// func Carousel_options.titleClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.titleClassName"] = p
// 	}
// }

// // 图片描述
// func Carousel_options.description(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.description"] = p
// 	}
// }

// // 图片描述类名
// func Carousel_options.descriptionClassName(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.descriptionClassName"] = p
// 	}
// }

// // HTML 自定义，同[Tpl](./tpl)一致
// func Carousel_options.html(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["options.html"] = p
// 	}
// }

// 自定义`schema`来展示数据
func Carousel_itemSchema(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["itemSchema"] = p
	}
}

// 是否自动轮播
func Carousel_auto(p bool) opt {
	return func(o map[string]interface{}) {
		o["auto"] = p
	}
}

// 切换动画间隔
func Carousel_interval(p string) opt {
	return func(o map[string]interface{}) {
		o["interval"] = p
	}
}

// 切换动画时长
func Carousel_duration(p string) opt {
	return func(o map[string]interface{}) {
		o["duration"] = p
	}
}

// 宽度
func Carousel_width(p string) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}

// 高度
func Carousel_height(p string) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// 显示左右箭头、底部圆点索引
func Carousel_controls(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["controls"] = p
	}
}

// 左右箭头、底部圆点索引颜色，默认`light`，另有`dark`模式
func Carousel_controlsTheme(p string) opt {
	return func(o map[string]interface{}) {
		o["controlsTheme"] = p
	}
}

// 切换动画效果，默认`fade`，另有`slide`模式
func Carousel_animation(p string) opt {
	return func(o map[string]interface{}) {
		o["animation"] = p
	}
}
