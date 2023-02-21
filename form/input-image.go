package form

// 图片
func InputImage(opts ...opt) map[string]interface{} {
	return newForm("input-image", opts...)
}

// 上传文件接口
func InputImage_receiver(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["receiver"] = p
	}
}

// 支持的图片类型格式，请配置此属性为图片后缀，例如`.jpg,.png`
func InputImage_accept(p string) opt {
	return func(o map[string]interface{}) {
		o["accept"] = p
	}
}

// 默认没有限制，当设置后，文件大小大于此值将不允许上传。单位为`B`
func InputImage_maxSize(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxSize"] = p
	}
}

// 默认没有限制，当设置后，一次只允许上传指定数量文件。
func InputImage_maxLength(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxLength"] = p
	}
}

// 是否多选。
func InputImage_multiple(p bool) opt {
	return func(o map[string]interface{}) {
		o["multiple"] = p
	}
}

// [拼接值](./options#%E6%8B%BC%E6%8E%A5%E5%80%BC-joinvalues)
func InputImage_joinValues(p bool) opt {
	return func(o map[string]interface{}) {
		o["joinValues"] = p
	}
}

// [提取值](./options#%E6%8F%90%E5%8F%96%E5%A4%9A%E9%80%89%E5%80%BC-extractvalue)
func InputImage_extractValue(p bool) opt {
	return func(o map[string]interface{}) {
		o["extractValue"] = p
	}
}

// [拼接符](./options#%E6%8B%BC%E6%8E%A5%E7%AC%A6-delimiter)
func InputImage_delimiter(p string) opt {
	return func(o map[string]interface{}) {
		o["delimiter"] = p
	}
}

// 否选择完就自动开始上传
func InputImage_autoUpload(p bool) opt {
	return func(o map[string]interface{}) {
		o["autoUpload"] = p
	}
}

// 隐藏上传按钮
func InputImage_hideUploadButton(p bool) opt {
	return func(o map[string]interface{}) {
		o["hideUploadButton"] = p
	}
}

// 如果你不想自己存储，则可以忽略此属性。
func InputImage_fileField(p string) opt {
	return func(o map[string]interface{}) {
		o["fileField"] = p
	}
}

// 用来设置是否支持裁剪。
func InputImage_crop(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["crop"] = p
	}
}

// // 裁剪比例。浮点型，默认 `1` 即 `1:1`，如果要设置 `16:9` 请设置 `1.7777777777777777` 即 `16 / 9`。。
// func InputImage_crop.aspectRatio(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["crop.aspectRatio"] = p
// 	}
// }

// // 裁剪时是否可旋转
// func InputImage_crop.rotatable(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["crop.rotatable"] = p
// 	}
// }

// // 裁剪时是否可缩放
// func InputImage_crop.scalable(p bool) opt {
// 	return func(o map[string]interface{}) {
// 		o["crop.scalable"] = p
// 	}
// }

// // 裁剪时的查看模式，0 是无限制
// func InputImage_crop.viewMode(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["crop.viewMode"] = p
// 	}
// }

// 裁剪文件格式
func InputImage_cropFormat(p string) opt {
	return func(o map[string]interface{}) {
		o["cropFormat"] = p
	}
}

// 裁剪文件格式的质量，用于 jpeg/webp，取值在 0 和 1 之间
func InputImage_cropQuality(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["cropQuality"] = p
	}
}

// 限制图片大小，超出不让上传。
func InputImage_limit(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["limit"] = p
	}
}

// 默认占位图地址
func InputImage_frameImage(p string) opt {
	return func(o map[string]interface{}) {
		o["frameImage"] = p
	}
}

// 是否开启固定尺寸,若开启，需同时设置 fixedSizeClassName
func InputImage_fixedSize(p bool) opt {
	return func(o map[string]interface{}) {
		o["fixedSize"] = p
	}
}

// 开启固定尺寸时，根据此值控制展示尺寸。例如`h-30`,即图片框高为 h-30,AMIS 将自动缩放比率设置默认图所占位置的宽度，最终上传图片根据此尺寸对应缩放。
func InputImage_fixedSizeClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["fixedSizeClassName"] = p
	}
}

// 表单反显时是否执行 autoFill
func InputImage_initAutoFill(p bool) opt {
	return func(o map[string]interface{}) {
		o["initAutoFill"] = p
	}
}

// 图片上传后是否进入裁剪模式
func InputImage_dropCrop(p bool) opt {
	return func(o map[string]interface{}) {
		o["dropCrop"] = p
	}
}

// 图片选择器初始化后是否立即进入裁剪模式
func InputImage_initCrop(p bool) opt {
	return func(o map[string]interface{}) {
		o["initCrop"] = p
	}
}

// 限制图片宽度。
func InputImage_width(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["width"] = p
	}
}

// 限制图片高度。
func InputImage_height(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["height"] = p
	}
}

// 限制图片最小宽度。
func InputImage_minWidth(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minWidth"] = p
	}
}

// 限制图片最小高度。
func InputImage_minHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["minHeight"] = p
	}
}

// 限制图片最大宽度。
func InputImage_maxWidth(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxWidth"] = p
	}
}

// 限制图片最大高度。
func InputImage_maxHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["maxHeight"] = p
	}
}

// 限制图片宽高比，格式为浮点型数字，默认 `1` 即 `1:1`，如果要设置 `16:9` 请设置 `1.7777777777777777` 即 `16 / 9`。 如果不想限制比率，请设置空字符串。
func InputImage_aspectRatio(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["aspectRatio"] = p
	}
}

// 上传文件值变化时触发(上传失败同样会触发)
func InputImage_change(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["change"] = p
	}
}

// 移除文件时触发
func InputImage_remove(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["remove"] = p
	}
}

// 上传文件失败时触发
func InputImage_fail(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["fail"] = p
	}
}
