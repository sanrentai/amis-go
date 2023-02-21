package form

// 富文本编辑器
func InputRichText(opts ...opt) map[string]interface{} {
	return newForm("input-rich-text", opts...)
}



// 是否保存为 ubb 格式
func InputRichText_saveAsUbb(p bool) opt {
	return func(o map[string]interface{}) {
		o["saveAsUbb"] = p
	}
} 

// 默认的图片保存 API
func InputRichText_receiver(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["receiver"] = p
	}
} 

// 默认的视频保存 API
func InputRichText_videoReceiver(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["videoReceiver"] = p
	}
} 

// 上传文件时的字段名
func InputRichText_fileField(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["fileField"] = p
	}
} 

// 框的大小，可设置为 `md` 或者 `lg`
func InputRichText_size(p string) opt {
	return func(o map[string]interface{}) {
		o["size"] = p
	}
} 

// 需要参考 [tinymce](https://www.tiny.cloud/docs/configure/integration-and-setup/) 或 [froala](https://www.froala.com/wysiwyg-editor/docs/options) 的文档
func InputRichText_options(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["options"] = p
	}
} 

// froala 专用，配置显示的按钮，tinymce 可以通过前面的 options 设置 [toolbar](https://www.tiny.cloud/docs/demo/custom-toolbar-button/) 字符串
func InputRichText_buttons(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["buttons"] = p
	}
} 