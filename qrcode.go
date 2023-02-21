package amis

// 二维码
func QRCode(opts ...opt) map[string]interface{} {
	return newCompent("qrcode", opts...)
}

// 外层 Dom 的类名
func QRCode_className(p string) opt {
	return func(o map[string]interface{}) {
		o["className"] = p
	}
}

// 二维码 SVG 的类名
func QRCode_qrcodeClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["qrcodeClassName"] = p
	}
}

// 二维码的宽高大小
func QRCode_codeSize(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["codeSize"] = p
	}
}

// 二维码背景色
func QRCode_backgroundColor(p string) opt {
	return func(o map[string]interface{}) {
		o["backgroundColor"] = p
	}
}

// 二维码前景色
func QRCode_foregroundColor(p string) opt {
	return func(o map[string]interface{}) {
		o["foregroundColor"] = p
	}
}

// 二维码复杂级别，有（'L' 'M' 'Q' 'H'）四种
func QRCode_level(p string) opt {
	return func(o map[string]interface{}) {
		o["level"] = p
	}
}

// 扫描二维码后显示的文本，如果要显示某个页面请输入完整 url（`"http://..."`或`"https://..."`开头），支持使用 [模板](./concepts/template)
func QRCode_value(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["value"] = p
	}
}

// QRCode 图片配置
func QRCode_imageSettings(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["imageSettings"] = p
	}
}

// // 图片链接地址
// func QRCode_imageSettings.src(p string) opt {
// 	return func(o map[string]interface{}) {
// 		o["imageSettings.src"] = p
// 	}
// }

// // 图片宽度
// func QRCode_imageSettings.width(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["imageSettings.width"] = p
// 	}
// }

// // 图片高度
// func QRCode_imageSettings.height(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["imageSettings.height"] = p
// 	}
// }

// // 图片水平方向偏移量
// func QRCode_imageSettings.x(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["imageSettings.x"] = p
// 	}
// }

// // 图片垂直方向偏移量
// func QRCode_imageSettings.y(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["imageSettings.y"] = p
// 	}
// }

// // 图片是否挖孔嵌入
// func QRCode_imageSettings.excavate(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["imageSettings.excavate"] = p
// 	}
// }
